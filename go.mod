module go.sancus.dev/config

go 1.16

replace go.sancus.dev/config/expand => ./expand

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/creasty/defaults v1.6.0
	github.com/hashicorp/hcl/v2 v2.16.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	go.sancus.dev/config/expand v0.0.0-00010101000000-000000000000
	go.sancus.dev/core v0.18.1
	golang.org/x/text v0.7.0 // indirect
	gopkg.in/dealancer/validate.v2 v2.1.0
	gopkg.in/yaml.v3 v3.0.1
)
