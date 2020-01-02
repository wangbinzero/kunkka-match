<img src="https://simpleicons.org/icons/go.svg" width="60px" height="60px"/>


# kunkka-match

高性能撮合引擎

### 连续竞价之价格判断

连续竞价要满足价格优先，时间优先原则
- 价格优先：买单价格较高者优先成交，卖单价格较低者优先成交
- 时间优先：买卖方向和价格相同的委托单，先上报的委托单会比后申报的委托单优先成交

- 买入价必须大于等于卖出价才能成交，当买入价等于买入价时，成交价就是买入或者卖出价。
- 当买入价大于卖出价时，则还需要参考前一笔成交价来确定最新成交价。假设买入价为 B 卖出价为S
前一笔成交价为P ，最新成交价为N  那么：
   - 如果 P >= B 则 N = B
   - 如果 P <= S 则 N = S
   - 如果 B > P > S 则 N = P 

### 中间件支持

- rabbitmq
- redis

### 运行

```bash
go run main.go
```
### 编译&运行

linux
```bash
GOOS=linux GOARCH=amd64 go build main.go
./main
```

windows
```bash
GOOS=windows GOARCH=amd64 go build main.go
```

### 开启交易标的引擎
POST http://127.0.0.1:8080/openMatching
```json
{
	"symbol":"btcusdt",  //交易标的
	"price":7200         //开盘价
}
```

### 关闭交易标的引擎
POST http://127.0.0.1:8080/closeMatching

```json

{
   "symbol":["btcusdt","ethusdt"]
}
```



```bash
-----------------------------------------------------------------
   __ __          __    __          __  ___     __      __
  / //_/_ _____  / /__ / /_____ _  /  |/  /__ _/ /_____/ /
 / ,< / // / _ \/  '_//  '_/ _  / / /|_/ / _  / __/ __/ _ \ 
/_/|_|\_,_/_//_/_/\_\/_/\_\\_,_/ /_/  /_/\_,_/\__/\__/_//_/

             Kunkka 高性能撮合引擎-v0.1  goLang
-----------------------------------------------------------------
[KUNKKA] –– 2020/01/02 11:50:13.398527 [INFO]交易标的引擎map初始化成功
[KUNKKA] –– 2020/01/02 11:50:13.398816 [INFO]缓存服务redis [127.0.0.1:6379] 连接成功 
[KUNKKA] –– 2020/01/02 11:50:13.399150 [INFO]订单簿初始化成功
[KUNKKA] –– 2020/01/02 11:50:13.399174 [INFO]撮合引擎 [btcusdt] 启动成功
[KUNKKA] –– 2020/01/02 11:50:13.399544 [INFO]缓存加载订单: 交易标的 [btcusdt] 订单 map[action:create orderId:1234568 orderType:limit price:1200 symbol:btcusdt timestamp:0.00000000000000000000000000000000000000000000000015779359203800869]
[KUNKKA] –– 2020/01/02 11:50:13.399736 [INFO]缓存加载订单: 交易标的 [btcusdt] 订单 map[action:create orderId:1234567 orderType:limit price:1200 symbol:btcusdt timestamp:0.0000000000000000000000000000000000000000000000001577935932283284]
[KUNKKA] –– 2020/01/02 11:50:14.401942 [INFO]服务端口: :8080

```