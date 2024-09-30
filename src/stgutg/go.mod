module stgutgp

go 1.20

replace (
	free5gclib => ../free5gclib
	tglib => ../tglib
)

require (
	free5gclib v0.0.0-00010101000000-000000000000
	github.com/ishidawataru/sctp v0.0.0-20210707070123-9a39160e9062
	gopkg.in/yaml.v2 v2.4.0
	tglib v0.0.0-00010101000000-000000000000
)

require (
	github.com/aead/cmac v0.0.0-20160719120800-7af84192f0b1 // indirect
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/calee0219/fatal v0.0.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/kr/text v0.1.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/wmnsk/milenage v1.2.1 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
