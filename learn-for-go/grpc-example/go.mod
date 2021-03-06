module github.com/Win-Man/sg-server

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20200327173247-9dae0f8f5775 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200326112834-f447254575fd // indirect
	google.golang.org/grpc v1.28.0
)

replace github.com/Win-Man/sg-server/protos => ./sg-server/protos
