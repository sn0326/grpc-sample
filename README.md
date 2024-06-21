# grpc-sample

## 事前準備

### protobufインストール

```bat
winget install --id=Google.Protobuf  -e
```

```bat
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 参考

- [Go言語gRPCを爆速で試す方法 \#grpc\-web \- Qiita](https://qiita.com/totoaoao/items/6bf533b6d2164b74ac09)