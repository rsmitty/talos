module github.com/talos-systems/talos

go 1.13

replace (
	github.com/Azure/go-autorest v10.8.1+incompatible => github.com/Azure/go-autorest/autorest v0.9.1
	github.com/docker/distribution v2.7.1+incompatible => github.com/docker/distribution v2.7.1-0.20190205005809-0d3efadf0154+incompatible
	github.com/talos-systems/talos/pkg/machinery => ./pkg/machinery
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Microsoft/hcsshim v0.8.9 // indirect
	github.com/Microsoft/hcsshim/test v0.0.0-20200831205110-d2cba219a8d7 // indirect
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2
	github.com/beevik/ntp v0.3.0
	github.com/containerd/cgroups v0.0.0-20200824123100-0b889c03f102
	github.com/containerd/containerd v1.4.1
	github.com/containerd/continuity v0.0.0-20200928162600-f2cc35102c2a // indirect
	github.com/containerd/cri v1.11.1
	github.com/containerd/go-cni v1.0.1
	github.com/containerd/ttrpc v1.0.2 // indirect
	github.com/containerd/typeurl v1.0.1
	github.com/containernetworking/cni v0.8.0
	github.com/containernetworking/plugins v0.8.7
	github.com/coreos/go-iptables v0.4.5
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.9.0
	github.com/firecracker-microvm/firecracker-go-sdk v0.21.0
	github.com/fullsailor/pkcs7 v0.0.0-20190404230743-d7302db945fa
	github.com/gizak/termui/v3 v3.1.0
	github.com/gogo/googleapis v1.4.0 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/hashicorp/go-getter v1.5.0
	github.com/hashicorp/go-multierror v1.1.0
	github.com/insomniacslk/dhcp v0.0.0-20200922210017-67c425063dca
	github.com/jsimonetti/rtnetlink v0.0.0-20201022080559-7cf95b6356e5
	github.com/kubernetes-sigs/bootkube v0.14.1-0.20200817205730-0b4482256ca1
	github.com/mattn/go-isatty v0.0.12
	github.com/mdlayher/genetlink v1.0.0
	github.com/mdlayher/netlink v1.1.1
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v1.0.0-rc92 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20200728170252-4d89ac9fbff6
	github.com/pin/tftp v2.1.0+incompatible
	github.com/prometheus/procfs v0.2.0
	github.com/rs/xid v1.2.1
	github.com/ryanuber/columnize v2.1.2+incompatible
	github.com/smira/go-xz v0.0.0-20150414201226-0c531f070014
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2
	github.com/talos-systems/bootkube-plugin v0.0.0-20200915135634-229d57e818f3
	github.com/talos-systems/crypto v0.2.1-0.20201112141136-12a489768a6b
	github.com/talos-systems/go-blockdevice v0.1.1-0.20201111103554-874213371a3f
	github.com/talos-systems/go-loadbalancer v0.1.0
	github.com/talos-systems/go-procfs v0.0.0-20200219015357-57c7311fdd45
	github.com/talos-systems/go-retry v0.1.1-0.20200922131245-752f081252cf
	github.com/talos-systems/go-smbios v0.0.0-20200807005123-80196199691e
	github.com/talos-systems/grpc-proxy v0.2.0
	github.com/talos-systems/net v0.2.0
	github.com/talos-systems/talos/pkg/machinery v0.0.0-20200818212414-6a7cc0264819
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae // indirect
	github.com/vmware-tanzu/sonobuoy v0.19.0
	github.com/vmware/vmw-guestinfo v0.0.0-20200218095840-687661b8bd8e
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200819165624-17cef6e3e9d5 // v3.4.10
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sync v0.0.0-20201008141435-b3e1573b7520
	golang.org/x/sys v0.0.0-20201018230417-eeed37f84f13
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	golang.org/x/tools v0.0.0-20201019175715-b894a3290fff // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/freddierice/go-losetup.v1 v1.0.0-20170407175016-fc9adea44124
	gopkg.in/fsnotify.v1 v1.4.7
	gopkg.in/yaml.v2 v2.2.8
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776
	k8s.io/api v0.19.3
	k8s.io/apimachinery v0.19.3
	k8s.io/apiserver v0.19.3
	k8s.io/client-go v0.19.3
	k8s.io/cri-api v0.19.3
	k8s.io/kubelet v0.19.3
)
