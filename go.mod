module stgutgmain

go 1.14

replace (
	free5gc/lib => ./src/free5gc/lib
	stgutg => ./src/stgutg
	tglib => ./src/tglib
)

require (
	free5gc/lib v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
	stgutg v0.0.0-00010101000000-000000000000
	tglib v0.0.0-00010101000000-000000000000
)
