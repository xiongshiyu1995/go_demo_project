#源镜像
FROM alpine:latest
#设置工作目录
WORKDIR /home/http/
#将服务器的go工程代码加入到docker容器中
COPY ./go-http .
#新增日志文件夹
RUN mkdir logs
#暴露端口
EXPOSE 80
#最终运行docker的命令
CMD nohup ./go-http > logs/app.log 2>&1
