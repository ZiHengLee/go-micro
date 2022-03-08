用于生成对应Go语言micro代码
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/micro/micro/v2/cmd/protoc-gen-micro

# 使用下面这个命令会生成user.pb.go/user.pb.micro.go的文件
protoc .proto --go_out=.
protoc --go_out=. --micro_out=. user.proto