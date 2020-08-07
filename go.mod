module github.com/meshplus/pier

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/cbergoon/merkletree v0.2.0
	github.com/fatih/color v1.9.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.6.3
	github.com/gobuffalo/packd v1.0.0
	github.com/gobuffalo/packr v1.30.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.3
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.3.0
	github.com/libp2p/go-libp2p-core v0.5.6
	github.com/meshplus/bitxhub-core v0.1.0-rc1.0.20200526060151-b0efad4a2046
	github.com/meshplus/bitxhub-kit v1.0.1-0.20200525112026-df2160653e23
	github.com/meshplus/bitxhub-model v1.0.0-rc4.0.20200731025300-2bb1717059e0
	github.com/meshplus/go-bitxhub-client v1.0.0-rc4.0.20200731031000-ec0387c42327
	github.com/meshplus/go-lightp2p v0.0.0-20200807082129-74180de79d19
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.2.2
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/goleveldb v1.0.1-0.20190923125748-758128399b1d
	github.com/tidwall/gjson v1.6.0
	github.com/urfave/cli v1.22.1
	golang.org/x/tools v0.0.0-20200102140908-9497f49d5709 // indirect
	google.golang.org/grpc v1.27.1
)

replace github.com/meshplus/go-lightp2p => ../go-lightp2p
