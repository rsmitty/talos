// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package packet

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/talos-systems/go-procfs/procfs"

	"github.com/talos-systems/talos/internal/app/machined/pkg/runtime"
	"github.com/talos-systems/talos/pkg/download"
	"github.com/talos-systems/talos/pkg/machinery/config/configloader"
	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1"
)

// Metadata holds packet metadata info.
type Metadata struct {
	Network        Network  `json:"network"`
	PrivateSubnets []string `json:"private_subnets"`
}

// Network holds network info from the packet metadata.
type Network struct {
	Bonding    Bonding     `json:"bonding"`
	Interfaces []Interface `json:"interfaces"`
	Addresses  []Address   `json:"addresses"`
}

// Bonding holds bonding info from the packet metadata.
type Bonding struct {
	Mode int `json:"mode"`
}

// Interface holds interface info from the packet metadata.
type Interface struct {
	Name string `json:"name"`
	Bond string `json:"bond"`
}

// Address holds address info from the packet metadata.
type Address struct {
	Public  bool   `json:"public"`
	Enabled bool   `json:"enabled"`
	CIDR    int    `json:"cidr"`
	Family  int    `json:"address_family"`
	Netmask string `json:"netmask"`
	Network string `json:"network"`
	Address string `json:"address"`
	Gateway string `json:"gateway"`
}

// BondMode is the list of ints for a bonding setup.
// nb: this is stolen from networkd code since it's not public in the machinery pkg.
//     used to translate the int returned from packet metadata to the string expected by networkd.
// https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/if_bonding.h
type BondMode uint8

// nolint: golint,stylecheck
const (
	BOND_MODE_ROUNDROBIN BondMode = iota
	BOND_MODE_ACTIVEBACKUP
	BOND_MODE_XOR
	BOND_MODE_BROADCAST
	BOND_MODE_8023AD
	BOND_MODE_TLB
	BOND_MODE_ALB
)

func (b BondMode) String() string {
	return [...]string{"balance-rr", "active-backup", "balance-xor", "broadcast", "802.3ad", "balance-tlb", "balance-alb"}[int(b)]
}

const (
	// PacketUserDataEndpoint is the local metadata endpoint for Packet.
	PacketUserDataEndpoint = "https://metadata.platformequinix.com/userdata"
	// PacketMetaDataEndpoint is the local endpoint for machine info like networking.
	PacketMetaDataEndpoint = "https://metadata.platformequinix.com/metadata"
)

// Packet is a discoverer for non-cloud environments.
type Packet struct{}

// Name implements the platform.Platform interface.
func (p *Packet) Name() string {
	return "packet"
}

// Configuration implements the platform.Platform interface.
// nolint: gocyclo
func (p *Packet) Configuration(ctx context.Context) ([]byte, error) {
	// Fetch and unmarshal both the talos machine config and the
	// metadata about the instance from packet's metadata server
	log.Printf("fetching machine config from: %q", PacketUserDataEndpoint)

	machineConfigDl, err := download.Download(ctx, PacketUserDataEndpoint)
	if err != nil {
		return nil, err
	}

	log.Printf("fetching equinix network config from: %q", PacketMetaDataEndpoint)

	metadataConfig, err := download.Download(ctx, PacketMetaDataEndpoint)
	if err != nil {
		return nil, err
	}

	var unmarshalledMetadataConfig Metadata
	if err = json.Unmarshal(metadataConfig, &unmarshalledMetadataConfig); err != nil {
		return nil, err
	}

	confProvider, err := configloader.NewFromBytes(machineConfigDl)
	if err != nil {
		return nil, err
	}

	var machineConfig *v1alpha1.Config
	machineConfig, ok := confProvider.(*v1alpha1.Config)
	if !ok {
		return nil, fmt.Errorf("unable to determine machine config type")
	}

	// translate the int returned from bond mode to the string needed by networkd
	bondMode := BondMode(uint8(unmarshalledMetadataConfig.Network.Bonding.Mode))

	// determine bond name and build list of interfaces enslaved by the bond
	// nb: currently only eth0 is supported.
	devicesInBond := []string{}
	bondName := ""

	for _, iface := range unmarshalledMetadataConfig.Network.Interfaces {
		if bondName != "" && iface.Bond != bondName {
			return nil, fmt.Errorf("encountered multiple bonds. this is unexpected in the equinix metal platform")
		}

		// TODO: allow for all ifaces to be in the bond once we figure out eth1 in packet.
		if iface.Name == "eth0" {
			devicesInBond = append(devicesInBond, iface.Name)
		}

		bondName = iface.Bond
	}

	// create multiple bond devices and add them to device list.
	// they will all get merged by networkd to configure the bond.
	packetDevices := []*v1alpha1.Device{}

	for _, addr := range unmarshalledMetadataConfig.Network.Addresses {
		bondDev := v1alpha1.Device{
			DeviceInterface: bondName,
			DeviceDHCP:      false,
			DeviceCIDR:      fmt.Sprintf("%s/%d", addr.Address, addr.CIDR),
			DeviceBond: &v1alpha1.Bond{
				BondMode:       bondMode.String(),
				BondInterfaces: devicesInBond,
			},
		}

		// "Public" interfaces get the default route
		if addr.Public {
			// TODO: suporrt ipv6 default route when we support them in networkd.
			if addr.Family == 4 {
				nw := "0.0.0.0/0"

				bondDev.DeviceRoutes = []*v1alpha1.Route{
					{
						RouteNetwork: nw,
						RouteGateway: addr.Gateway,
					},
				}
			}
		} else {
			// for "Private" interfaces, we add a route that goes out the gateway for the private subnets.
			for _, privSubnet := range unmarshalledMetadataConfig.PrivateSubnets {
				privRoute := &v1alpha1.Route{
					RouteNetwork: privSubnet,
					RouteGateway: addr.Gateway,
				}

				bondDev.DeviceRoutes = append(bondDev.DeviceRoutes, privRoute)
			}
		}

		packetDevices = append(packetDevices, &bondDev)
	}

	machineConfig.MachineConfig.MachineNetwork.NetworkInterfaces = append(
		machineConfig.MachineConfig.MachineNetwork.NetworkInterfaces,
		packetDevices...,
	)

	return confProvider.Bytes()
}

// Mode implements the platform.Platform interface.
func (p *Packet) Mode() runtime.Mode {
	return runtime.ModeMetal
}

// Hostname implements the platform.Platform interface.
func (p *Packet) Hostname(context.Context) (hostname []byte, err error) {
	return nil, nil
}

// ExternalIPs implements the runtime.Platform interface.
func (p *Packet) ExternalIPs(context.Context) (addrs []net.IP, err error) {
	return addrs, err
}

// KernelArgs implements the runtime.Platform interface.
func (p *Packet) KernelArgs() procfs.Parameters {
	return []*procfs.Parameter{
		procfs.NewParameter("console").Append("ttyS1,115200n8"),
	}
}
