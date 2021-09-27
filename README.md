# Go GRPC example

## Generating Pb files
````shell
protoc -I proto --go_out=plugins=grpc:internal/net/grpc proto/wishlist.proto
````

## Sources

* [Introducción a gRPC en Go - Friends of Go](https://blog.friendsofgo.tech/posts/introduccion-a-grpc/)
* [Protocol Buffers - Google](https://developers.google.com/protocol-buffers)