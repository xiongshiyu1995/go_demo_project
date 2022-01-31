# go_demo_project

# 模块3作业
docker build -t xiongshiyu/simple-http:v0.1.0 .  
docker run --name simple-http -p 80:80 -d xiongshiyu/simple-http:v0.1.0  
docker exec -it $containerId /bin/sh  
docker login  
docker push xiongshiyu/simple-http:v0.1.0  