# 搭建prometheus
### 1. Dockerfile ###
> zz/random 
    
````
cd client/client_golang/examples/random
docker build -t zz/random .
````

> zz/exe
````
cd notify
docker build -t zz/exe .
````
***
### 2. 构建docker-compose ### 
````
docker-compose up -d
````
***
### 3. 测试 ### 
````
cat <<EOF | curl --data-binary @- http://192.168.100.102:9091/metrics/job/cqh/instance/test
# 锻炼场所价格
muscle_metric{label="gym"} 8800
# 三大项数据 kg
bench_press 24
dead_lift 177
deep_squal 37
EOF
````