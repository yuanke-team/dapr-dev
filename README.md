# 官方安装入门教程
https://docs.microsoft.com/zh-cn/dotnet/architecture/dapr-for-net-developers/getting-started

# dapr install 
```
wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash

https://github.com/dapr/cli/releases/download/v1.5.1/dapr_linux_amd64.tar.gz 
cp to /bin
```
# 启动
./web-port/run_web.sh
./dapr-api-go/run_api_go.sh

# dapr 启动模式 
```
// 依赖Docker的模式 ，这个比较方便启动
dapr init  

// 不依赖Docker的模式 ，需要自己启动 ~/.dapr/bin/placement 
dapr init --slim

nohup  ~/.dapr/bin/placement &
```
# 配置consul server集群 参考：http://www.hdget.com/dapr-consul-config/
```
<!-- 下载连接：https://releases.hashicorp.com/consul -->
<!-- onsul集群的节点(raft共识) -bootstrap-expect 集群节点数 -bind 本机ip -join 要加入的初始节点ip-->
nohup ./consul agent -server -bootstrap-expect 2 -data-dir /tmp/consul -ui -client 0.0.0.0 -bind 192.168.0.5 -join 192.168.0.5 > /var/log/consul.log 2>&1 &
   
nohup ./consul agent -server -ui -data-dir /tmp/consul -client 0.0.0.0 -bind 192.168.0.4 -join 192.168.0.5 > /var/log/consul.log 2>&1 &

   
```
# dapr客户端()配置
###  nano ~/.dapr/config.yaml
```
apiVersion: dapr.io/v1alpha1
kind: Configuration
metadata:
  name: daprConfig
spec:
  nameResolution:
    component: "consul"
    configuration:
      client:
        address: "192.168.0.5:8500"
      selfRegister: true

OR

apiVersion: dapr.io/v1alpha1
kind: Configuration
metadata:
  name: appconfig
spec:
  nameResolution:
    component: "consul"
    configuration:
      selfRegister: false


nohup ./consul agent -data-dir /tmp/consul -ui -client 0.0.0.0 -bind 192.168.0.7 -join 192.168.0.4 > /var/log/consul.log 2>&1 &


```

# dapr服务端(需要注册发现的服务放这)配置
```
apiVersion: dapr.io/v1alpha1
kind: Configuration
metadata:
  name: appconfig
spec:
  nameResolution:
    component: "consul"
    configuration:
      selfRegister: true


OR 

nohup ./consul agent -data-dir /tmp/consul -ui -client 0.0.0.0 -bind 192.168.0.6 -join 192.168.0.5 > /var/log/consul.log 2>&1 &

```

# 注册点
```
除了consul bootstrap-expect 节点不可以运行dapr程序外其它的都可以
```