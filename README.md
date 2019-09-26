# golang版本控制台命令
golang实现的控制台启动脚本
这用于jar包项目启动脚本
在windows或者linux的启动脚本
## 配置文件
app.conf

# 执行的命令
run = 控制台执行的命令
如这里用来启动jar包:
run = java -jar guns-rest-0.0.1.jar

# 输出的日志
log = all
file:输出文件，console:只输出到控制台，all:文件和控制台都输出

# 执行的平台，windows或者linux
profile = windows

## 项目编译打包
linux:
set GOARCH=amd64
set GOOS=linux
go build main.go

linux:
set GOARCH=amd64
set GOOS=windows
go build main.go
上传到linux服务器后
chmod 777 main

# 注意事项
编译的后的可执行文件main必须和conf文件夹在同一路径下
即:可执行文件必须和conf文件夹一起打包上传服务器