module nodeid-server

require (
	github.com/google/gops v0.3.5
	github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1 // indirect
	github.com/nkawa/synerex_core_test/nodeapi v0.0.0
	google.golang.org/grpc v1.23.0
)

replace github.com/nkawa/synerex_core_test/nodeapi => ../nodeapi
