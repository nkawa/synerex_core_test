module synerex_server

go 1.12

require (
	github.com/nkawa/synerex_core_test/api v0.0.0-00010101000000-000000000000
	github.com/nkawa/synerex_core_test/sxutil v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/grpc v1.23.0
)

replace (
	github.com/nkawa/synerex_core_test/api => ../api
	github.com/nkawa/synerex_core_test/sxutil => ../sxutil
)
