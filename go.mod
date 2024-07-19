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
	github.com/google/gopacket v1.1.19 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/ishidawataru/sctp v0.0.0-20210707070123-9a39160e9062 // indirect
	github.com/j-keck/arping v1.0.3 // indirect
	github.com/libp2p/go-netroute v0.2.1 // indirect
	github.com/prometheus-community/pro-bing v0.3.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/wmnsk/milenage v1.2.1 // indirect
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.14.1-0.20231108175955-e4099bfacb8c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
