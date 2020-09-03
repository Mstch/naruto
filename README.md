## Naruto

一个说到做到的牛逼KV   

![](doc/858974c8bba3e7cc2e7d8989c8203522.gif)

```
├── conf                    
├── doc                     
├── helper------------------与raft逻辑无关的模块,抽离出来防止代码太乱
│   ├── db------------------封装了一下pebble
│   │   └── test-data       
│   ├── event --------------观察者模式,估计用不上
│   ├── logger
│   ├── member--------------集群成员管理
│   ├── rpc-----------------强依赖protobuf的rpc接口定义
│   │   └── stupid----------rpc的一个实现,具有池化的buffer,目前只实现了异步调用
│   ├── timer---------------具有随机定时功能的定时器
│   └── util----------------目前只写了几个[]byte与整数来回转换的逻辑
├── raft---------------------raft逻辑
│   ├── candidateHandler.go
│   ├── leaderHandler.go
│   ├── followerHandler.go---三个角色下对rpc,timer timeout事件的handler实现
│   ├── handler.go-----------handler注册与通用handler(term,commit index)实现
│   ├── client.go------------rpcclient注册   
│   ├── kv.go----------------kv存储实现,依赖helper里的db
│   ├── log.go---------------log存储实现,依赖helper里的db
│   ├── stat.go--------------节点状态实现,依赖helper里的db
│   ├── msg------------------raft所需的rpc消息体的定义
│   │   ├── msg.pb.go
│   │   ├── msg.proto
│   └── raft.go--------------启动器            
└── startup.sh---------------docker测试启动器
```

有状态startup有依赖的顺序:
member->server
statDB->timer
