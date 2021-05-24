module github.com/talos-systems/talos/pkg/machinery

go 1.16

// forked go-yaml that introduces RawYAML interface, which can be used to populate YAML fields using bytes
// which are then encoded as a valid YAML blocks with proper indentiation
replace gopkg.in/yaml.v3 => github.com/unix4ever/yaml v0.0.0-20210315173758-8fb30b8e5a5b

require (
	github.com/AlekSi/pointer v1.1.0
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef
	github.com/containerd/go-cni v1.0.2
	github.com/containernetworking/cni v0.8.1 // indirect
	github.com/cosi-project/runtime v0.0.0-20210409233936-10d6103c19ab
	github.com/dustin/go-humanize v1.0.0
	github.com/evanphx/json-patch v4.9.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-multierror v1.1.1
	github.com/onsi/ginkgo v1.15.0 // indirect
	github.com/onsi/gomega v1.10.3 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20200929063507-e6143ca7d51d
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/talos-systems/crypto v0.2.1-0.20210427105118-4f80b976b640
	github.com/talos-systems/go-blockdevice v0.2.1-0.20210510233948-1292574643e0
	github.com/talos-systems/net v0.2.1-0.20210212213224-05190541b0fa
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20210324051608-47abb6519492 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210302174412-5ede27ff9881
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)
