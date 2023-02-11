module go.sancus.dev/config/flags/cobra

go 1.16

replace (
	go.sancus.dev/config => ../..
	go.sancus.dev/config/expand => ../../expand
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	go.sancus.dev/config v0.10.3
)
