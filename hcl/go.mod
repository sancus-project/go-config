module go.sancus.dev/config/hcl

go 1.16

replace (
	go.sancus.dev/config => ../
	go.sancus.dev/config/expand => ../expand
)

require (
	github.com/hashicorp/hcl/v2 v2.16.0
	go.sancus.dev/config v0.10.3
	go.sancus.dev/config/expand v0.0.0-00010101000000-000000000000
	go.sancus.dev/core v0.18.1
	gopkg.in/dealancer/validate.v2 v2.1.0
)

require (
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	golang.org/x/text v0.7.0 // indirect
)
