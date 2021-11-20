module github.com/cloudnativedesign/contentmanager

go 1.17

require (
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/sys v0.0.0-20211109065445-02f5c0300f6e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)

replace github.com/cloudnativedesign/contentmanager/server v0.0.0 => ./server

replace github.com/cloudnativedesign/contentmanager/server/models v0.0.0-unpublished => ./server/models

replace github.com/cloudnativedesign/contentmanager/server/controllers v0.0.0-unpublished => ./server/controllers
