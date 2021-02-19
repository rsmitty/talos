---
title: "Disk Encryption"
description: "Guide on using system disk encryption"
---

It is possible to enable system disk partitions encryption on the OS level.
Only STATE and EPHEMERAL partitions can be encrypted at the moment.
STATE contains the most sensitive node data: secrets and certs.
EPHEMERAL partition can contain some sensitive workload data.

Data is encrypted using LUKS2, which is provided by the Linux kernel modules and `cryptsetup` utility.

System runs additional steps when encryption is enabled.
If the disk encryption is enabled for the STATE partition, the system will:

- save STATE encryption config as JSON in the META partition.
- before mounting the STATE partition, load encryption configs either from the machine config or from the META partition,
  the machine config is always preferred over the META one.
- before mounting the STATE partition, format and encrypt it, but only if it is empty and has no filesystem.

If the disk encryption is enabled for the EPHEMERAL partition, the system will:

- get the encryption config from the machine config.
- before mounting the EPHEMERAL partition, encrypt and format it, but only if it is empty and has not filesystem.

## Configuration

Right now system partitions encryption is disabled by default.
To enable disk encryption you should modify the machine configuration.
The machine config has the appropriate sections for STATE and EPHEMERAL partitions:

```yaml
machine:
  ...
  systemDisksEncryption:
    ephemeral:
      keys:
        - nodeID:
          keySlot: 0
    state:
      keys:
        - nodeID:
          keySlot: 0
```

### Encryption Keys

> Note: LUKS2 docs calls that keys, but in reality it is a passphrase.
> When this passphrase is added, LUKS2 runs argon2 to create an actual key from that passphrase.

LUKS2 supports up to 32 encryption keys and it is possible to specify all of them in the machine configuration.
Talos always tries to sync the keys list defined in the machine config with the actual keys defined for the LUKS2 partition.
So if you update the keys list you should have at least one key that is not changed to be used for keys management.

When you define a key you should specify the key kind and the `keySlot`:

```yaml
machine:
  ...
  state:
    keys:
      - nodeID: # key kind
        keySlot: 1

  ephemeral:
    keys:
      - static:
          passphrase: supersecret
        keySlot: 0
```

Take a note that keys order does not play any role on which key slot is used.
Every key must always have a slot defined.

### Encryption Key Kinds

Talos supports two kinds of keys for now:

- `nodeID` which is generated using the node UUID and the partition label (note that if the node UUID is not really random it will fail the entropy check).
- `static` which you define right in the configuration.

> Note: Use static keys only if your STATE partition is encrypted and only for the EPHEMERAL partition.
> For the STATE partition it will be stored in the META partition, which is not encrypted.

### Key Rotation

It is necessary to do `talosctl apply-config` couple of times to rotate the keys.
The idea is to always keep using at least a single working key and change everything around it.

So, for example, first add a new key:

```yaml
machine:
  ...
  ephemeral:
    keys:
      - static:
          passphrase: oldkey
        keySlot: 0
      - static:
          passphrase: newkey
        keySlot: 1
  ...
```

Run:

```bash
talosctl apply-config -n <node> -f config.yaml
```

Then remove the old key:

```yaml
machine:
  ...
  ephemeral:
    keys:
      - static:
          passphrase: newkey
        keySlot: 1
  ...
```

Run:

```bash
talosctl apply-config -n <node> -f config.yaml
```

## Going from Unencrypted to Encrypted and Vice Versa

### Ephemeral Partition

There is none in place encryption support for the partitions right now, so to avoid losing any data
only the empty partitions can be encrypted.
So the key is the explicit wipe call.

That's why migration from unencrypted to encrypted needs some additional handling, which is:

- `apply-config` should be called with `--on-reboot` flag.
- partition should be wiped after `apply-config`, but before the reboot.

Edit your machine config and add the encryption configuration:

```bash
vim config.yaml
```

Apply the configuration with `--on-reboot` flag:

```bash
talosctl apply-config -f config.yaml -n <node ip> --on-reboot
```

Wipe the partition you're going to encrypt:

```bash
talosctl reset --system-labels-to-wipe EPHEMERAL -n <none> --reboot=true
```

That's it, after you run the last command, the partition will be wiped and the node will reboot.
During the next boot the system will encrypt the partition.

### State Partition

Calling wipe against the STATE partition will make the node lose the config, so the previous flow is not going to work.

The flow should be to first wipe the STATE partition:

```bash
talosctl reset  --system-labels-to-wipe STATE -n <node ip> --reboot=true
```

Node will enter into maintenance mode, then run `apply-config` with `--insecure` flag:

```bash
talosctl apply-config --insecure -n <node ip> -f config.yaml
```

After installation is complete the node should encrypt the STATE partition.
