module tglib

go 1.20

replace free5gclib => ../free5gclib

require (
	free5gclib v0.0.0-00010101000000-000000000000
	github.com/calee0219/fatal v0.0.1
	github.com/ishidawataru/sctp v0.0.0-20210707070123-9a39160e9062
	github.com/wmnsk/milenage v1.2.1
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

require (
	github.com/aead/cmac v0.0.0-20160719120800-7af84192f0b1 // indirect
	github.com/antonfisher/nested-logrus-formatter v1.3.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
)
