module stgutgmain

go 1.21.4

toolchain go1.21.6

replace (
	free5gclib => ./src/free5gclib
	stgutg => ./src/stgutg
	tglib => ./src/tglib
)

require (
	github.com/Rotchamar/xdp_gtp v0.2.1
	github.com/cilium/ebpf v0.12.3
	stgutg v0.0.0-00010101000000-000000000000
	tglib v0.0.0-00010101000000-000000000000
)

require (
	free5gclib v0.0.0-00010101000000-000000000000 // indirect
	github.com/aead/cmac v0.0.0-20160719120800-7af84192f0b1 // indirect
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/calee0219/fatal v0.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/free5gc/aper v1.0.2 // indirect
	github.com/free5gc/logger_conf v1.0.0 // indirect
	github.com/free5gc/logger_util v1.0.0 // indirect
	github.com/free5gc/nas v1.0.1 // indirect
	github.com/free5gc/ngap v1.0.4 // indirect
	github.com/free5gc/openapi v1.0.3 // indirect
	github.com/free5gc/path_util v1.0.0 // indirect
	github.com/free5gc/util_3gpp v1.0.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.1 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/ishidawataru/sctp v0.0.0-20210707070123-9a39160e9062 // indirect
	github.com/j-keck/arping v1.0.3 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/libp2p/go-netroute v0.2.1 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/prometheus-community/pro-bing v0.3.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/ugorji/go/codec v1.2.1 // indirect
	github.com/wmnsk/milenage v1.2.1 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.14.1-0.20231108175955-e4099bfacb8c // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
