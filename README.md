# golang版本控制台命令
golang实现的控制台启动脚本
这用于jar包项目启动脚本
在windows或者linux的启动脚本
## 配置文件
app.conf

# 执行的命令
run = java -jar guns-rest-0.0.1.jar

# 输出的日志
log = all
file:输出文件，console:只输出到控制台，all:文件和控制台都输出

# 执行的平台，windows或者linux
profile = windows