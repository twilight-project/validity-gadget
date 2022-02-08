module github.com/twilight-project/validity-gadget/forkoracle-go

go 1.17

require github.com/ybbus/jsonrpc/v2 v2.1.6

require (
	github.com/onsi/gomega v1.13.0 // indirect
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
)

replace github.com/twilight-project/validity-gadget-chain => ../validity-gadget-chain

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/tendermint/starport => ../starport
