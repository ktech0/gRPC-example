### 最基本的 gRPC 负载均衡用法，并引入 ETCD 用作服务发现。

### gRPC 负载均衡一般思路

- 根据服务 key 前缀从 ETCD 获取负载均衡的服务地址列表。
- 使用 gRPC 自带的客户端负载均衡方案，在服务地址列表里选取其中一个进行调用。
- 借助 ETCD 的 Watch 功能监测服务的变化，实时更新服务地址列表（动态维护 Build() 中的地址列表）。

#### 参考：https://www.bilibili.com/read/cv6653581