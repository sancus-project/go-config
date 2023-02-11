module go.sancus.dev/config

go 1.16

replace go.sancus.dev/config/expand => ./expand

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/creasty/defaults v1.6.0
	go.sancus.dev/config/expand v0.0.0-00010101000000-000000000000
	go.sancus.dev/core v0.18.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1
)
