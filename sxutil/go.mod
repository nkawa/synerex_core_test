module github.com/nkawa/synerex_core_test/sxutil

require (
	cloud.google.com/go v0.34.0 // indirect
	github.com/bwmarrin/snowflake v0.0.0-20180412010544-68117e6bbede
	github.com/golang/lint v0.0.0-20181217174547-8f45f776aaf1 // indirect
	github.com/golang/mock v1.2.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/nkawa/synerex_core_test/api v0.0.0
	github.com/nkawa/synerex_core_test/nodeapi v0.0.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/lint v0.0.0-20181217174547-8f45f776aaf1 // indirect
	golang.org/x/net v0.0.0-20181220203305-927f97764cc3 // indirect
	golang.org/x/oauth2 v0.0.0-20181203162652-d668ce993890 // indirect
	golang.org/x/sync v0.0.0-20181221193216-37e7f081c4d4 // indirect
	golang.org/x/sys v0.0.0-20190107070147-cb59ee366067 // indirect
	golang.org/x/tools v0.0.0-20190107030206-d345f29b0d7b // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20181221175505-bd9b4fb69e2f // indirect
	google.golang.org/grpc v1.23.0
	honnef.co/go/tools v0.0.0-20190106161140-3f1c8253044a // indirect
)

replace (
	github.com/nkawa/synerex_core_test/api => ../api
	github.com/nkawa/synerex_core_test/nodeapi => ../nodeapi
	github.com/nkawa/synerex_core_test/sxutil => ../sxutil
)
