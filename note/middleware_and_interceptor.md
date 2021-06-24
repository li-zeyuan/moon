# middleware and interceptor
## middleware
- http服务中间件
- 在请求处理业务handler前(后)，追加一些通用的逻辑处理

## interceptor
- grpc服务拦截器
- 类似与http服务中间件
- 分UnaryInterceptor，StreamInterceptor

## UnaryInterceptor
- 客户端、服务端一元（单向调用）
- 调用过程：
    - 预处理(pre-processing)
    - 调用RPC方法(invoking RPC method)
    - 后处理(post-processing)

## StreamInterceptor
- 流拦截器，客户端、服务端双向调用
- 调用过程：
    - 1、预处理
    - 2、调用SendMsg、RecvMsg
    - 3、后处理

## 流程
```sequence
client->user.logic.http: request
Note over  user.logic.http: RequestIdMiddleware：\n1、生成request_id，\n 2、set to content、response header
Note over  user.logic.http: InfraMiddleware：\n1、生成infra(db/logger)，\n2、set to content
user.logic.http->user.db.rpc: RequestIDClientInterceptor：\n1、在context中加入MD，\n2、在MD中加入 reqeust_id
Note over user.db.rpc: RequestIDUnaryServerInterceptor：\n1、从context中获取MD\n2、从MD中获取request_id
Note over user.db.rpc: InfraUnaryServerInterceptor：
user.db.rpc-->user.logic.http: response
user.logic.http-->client: response
```



## 参考
- gRPC interceptor 接受: https://grpc.io/blog/grpc-web-interceptor/
- gRPC传递request_id: https://xuanwo.io/2019/03/10/request-id-in-grpc/