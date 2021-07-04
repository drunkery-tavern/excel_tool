#!/bin/sh
#green='\e[92m'     #绿
blue='\e[34m'      #蓝
red='\e[31m'       #红
light_red='\e[91m' #浅红
res='\e[0m'
# 进程名
process="./bin/blog_go"
#文件目录
file_dir="./bin/blog_go"
#日志文件
log="blog_go.log"

setBuildEnv() {
  go env -w GOPROXY=https://goproxy.io,direct
  go env -w CGO_ENABLED=0
  go env -w GOOS=linux
  go env -w GOARCH=amd64
}

echo -e "${blue}build project... ${res}\n"
setBuildEnv
go mod tidy 2>&1
go build -o $process 2>&1
if [ $? -ne 0 ]; then
  echo -e "${red}build fail ${res}\n"
  exit 1
fi
echo -e "${blue}build successfully ${res}\n"
chmod +x ${file_dir}

# 获取进程ID
PID=$(ps -ef | grep $process | grep -v grep | awk '{print $2}')
echo -e "${blue}start deployment ${res}\n"
if [ -n "$PID" ]; then
  echo -e "${light_red}$process is running in $PID... ${res}\n"
  kill -9 "$PID"
  echo -e "${red}kill process $PID${res}\n"
else
  echo -e "${blue}$process is no exist...${res}\n"
fi
nohup $file_dir >$log 2>&1 &
echo -e "${blue}deployment successfully, process pid is:$(ps -ef | grep $process | grep -v grep | awk '{print $2}')${res}\n"