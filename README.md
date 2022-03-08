# go-micro-study
安装jaeger组件
 docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
http://localhost:16686/search查询界面

docker pull consul:1.6.0
docker run -d -p 8500:8500 consul:1.6.0
127.0.0.1:8500

微服务熔断观测面板安装
docker pull mlabouardy/hystrix-dashboard
docker run --name hystrix-dashboard -d -p 9002:9002 mlabouardy/hystrix-dashboard:latest
protoc .proto --go_out=.
protoc --go_out=. --micro_out=. user.proto


Docker-Compose安装使用
首先，使用Dockerfile定义应用程序的环境
其次，使用docker-compose.yml定义构成应用程序的服务
最后，执行docker-compose up命令来启动并运行整个应用程序
安装docker-compose

curl -L https://github.com/docker/compose/releases/download/1.26.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

chmod 777 /usr/local/bin/docker-compose

docker-compose --version
