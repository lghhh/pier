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
	github.com/libp2p/go-libp2p-core v0.3.0
	github.com/meshplus/bitxhub-core v0.1.0-rc1.0.20200526060151-b0efad4a2046
	github.com/meshplus/bitxhub-kit v1.0.1-0.20200525112026-df2160653e23
	github.com/meshplus/bitxhub-model v1.0.0-rc4.0.20200608065824-2fbc63639e92
	github.com/meshplus/go-bitxhub-client v1.0.0-rc4.0.20200509065005-851bf8c357e4
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-multiaddr v0.2.0
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.1-0.20190923125748-758128399b1d
	github.com/tidwall/gjson v1.6.0
	github.com/urfave/cli v1.22.1
	golang.org/x/sys v0.0.0-20200509044756-6aff5f38e54f // indirect
)

replace github.com/hyperledger/fabric => ../../hyperledger/fabric

replace github.com/meshplus/bitxhub-core => ../bitxhub-core

replace github.com/hyperledger/fabric-protos-go => github.com/hyperledger/fabric-protos-go v0.0.0-20200330074707-cfe579e86986

replace go.uber.org/zap => go.uber.org/zap v1.9.1

replace github.com/sykesm/zap-logfmt => github.com/sykesm/zap-logfmt v0.0.1

replace golang.org/x/net => golang.org/x/net v0.0.0-20200202094626-16171245cfb2

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20200509044756-6aff5f38e54f

replace github.com/sirupsen/logrus => github.com/sirupsen/logrus v1.5.0

replace google.golang.org/grpc => google.golang.org/grpc v1.27.1

replace github.com/mholt/archiver => github.com/mholt/archiver v0.0.0-20180417220235-e4ef56d48eb0

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

replace github.com/spf13/pflag => github.com/spf13/pflag v1.0.5

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.2

replace golang.org/x/text => golang.org/x/text v0.3.0

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.7
