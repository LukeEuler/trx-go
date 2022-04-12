# info

[tron](https://github.com/tronprotocol/protocol) 是直接copy的. `git submodule add` 有点大
commit: bf72b12fcb023be19773f4453c6a59189771ed3e
tag: GreatVoyage-v4.3.0

[googleapis](https://github.com/googleapis/googleapis) 是部分copy. 同上
commit: fb51f3cd6019088e76c88257a343bce8899d818f

bash plugin/trx/proto/gen-proto.sh
然后手动稍作修改. 如整理文件, 整理依赖


# env

```
// brew install protobuf
protoc --version
libprotoc 3.19.3

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

# ...

md, grpc 就是麻烦. 尤其是有这种自己不维护一个纯净sdk的, 还 tmd 弄得不干不净的