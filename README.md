#tfserving-r2.8 客户端grpc使用@go1.20

#使用方式

```
#cd /opt/
#git clone -b v1.0.0 https://github.com/howge/tfs28.git
#mkdir test && cd test
#go mod init test
#go work init
#go work use .
#go work use ../tfs28
#vim main.go
```
#内容如下:
```
package main

import (
	"fmt"

	config "github.com/howge/tfs28/config"
)

#测试文件是否能正常使用
func main() {
	c := config.FileSystemStoragePathSourceConfig{}
	fmt.Println(c)
}

```


- 根据tfserving描述，只需要将核心proto文件通过protoc 转成go文件即可，可以分为如下几步：

- 克隆分支将核心文件转为*.pb.go
  
```
git clone -b r2.8 https://github.com/tensorflow/serving.git
git clone -b r2.8 https://github.com/tensorflow
```

- 生成tf库文件

```
find ./tensorflow/tensorflow/core/framework  -name "*.proto" -exec protoc -I serving -I tensorflow  --go_out=plugins=grpc:vendor  {} \;
find ./tensorflow/tensorflow/core/example  -name "*.proto" -exec protoc -I serving -I tensorflow  --go_out=plugins=grpc:vendor  {} \;
find ./tensorflow/tensorflow/core/protobuf  -name "*.proto" -exec protoc -I serving -I tensorflow  --go_out=plugins=grpc:vendor  {} \;
```

- 生成tfserving 库文件，在生成库文件之前需要给proto文件加一个go_package 选项，格式参考如下:
```
find ./tensorflow_serving/config -name "*.proto" -exec  sed  -in '/syntax = "proto3";/a\option go_package = "github.com/howge/tfs28/config";' {} \;
find ./tensorflow_serving/apis -name "*.proto" -exec  sed  -in '/syntax = "proto3";/a\option go_package = "github.com/howge/tfs28/apis";' {} \;
find ./tensorflow_serving/resources -name "*.proto" -exec  sed  -in '/syntax = "proto3";/a\option go_package = "github.com/howge/tfs28/resources";' {} \;
find ./tensorflow_serving/sources  -name "*.proto" -exec  sed  -in '/syntax = "proto3";/a\option go_package = "github.com/howge/tfs28/sources";' {} \;
  ```

- 生成tfserving 库文件

```
find ./serving/tensorflow_serving/ -name "*.proto" -exec protoc -I serving -I tensorflow  --go_out=plugins=grpc:vendor  {} \;
find ./serving/tensorflow_serving/ -name "*.proto" -exec protoc -I serving -I tensorflow  --go_out=plugins=grpc:vendor  {} \;
find ./serving/tensorflow_serving/ -name "*.proto" -exec protoc -I serving -I tensorflow  --go_out=plugins=grpc:vendor  {} \;
```

# 总结

- 原理其实很简单，就是根据proto文件生成go文件，核心就是解决依赖关系，找了一圈没有合适的库(版本不匹配)，只能手搓(踩坑），具体可以看go.mod 文件，其中对tensorflow的依赖改成本地引用
  
- go的mod管理,会报下面的错误，未能解决

```
#go get github.com/howge/tfs28@v1.0.0
go: github.com/howge/tfs28@v1.0.0: parsing go.mod:
      module declares its path as: tfs28
              but was required as: github.com/howge/tfs28
```


 
