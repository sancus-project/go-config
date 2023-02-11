module go.sancus.dev/config/yaml

go 1.16

replace (
	go.sancus.dev/config => ../
	go.sancus.dev/config/expand => ../expand
)

require (
	go.sancus.dev/config v0.10.3
	go.sancus.dev/core v0.18.1
	gopkg.in/yaml.v3 v3.0.1
)

require go.sancus.dev/config/expand v0.1.0
