package arch

type archType string

const (
	JSONRPC = archType("JSON-RPC")
	REST    = archType("REST")
)
