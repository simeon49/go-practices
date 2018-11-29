##Protocol Buffers

参考: [https://developers.google.com/protocol-buffers/docs/gotutorial](https://developers.google.com/protocol-buffers/docs/gotutorial)

####安装protocol-buffers
```bash
# 下载 https://github.com/protocolbuffers/protobuf/releases/
# 解压
$ ./configure
$ make; make install;
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

####编译 proto文件
```bash
$ protoc ./addressbook.proto --go_out=./pb
```
