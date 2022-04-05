## project layout
```shell
├── cmd                    application
├── configuration          服務有哪些控制項
├── deployment             服務部署資訊包括 Config file 和 dockerfile 但目前沒有在使用
├── docker-compose.yml     本地需要的一些服務
├── docs                   文件擺放地方 ( swagger or graphql 文件)
│   └── graph              GraphQL 如果要更改 schema 會到這個資料夾更改 然後 make gen.graphql
│       └── schema
│           ├── app
│           └── platform
├── internal               一些各種第三方元件的包裝
├── makefile               基本上所有操作指令都會使用 makefile
├── pkg                    服務有的模組
│   ├── converter          轉換模組 兩方的結構皆不存在 pkg 底下, 拆開後以後獨立出去會較簡單
│   │   └── auth.go
│   ├── delivery           傳輸層, 可分為 gRPC, RESTful, GraphQL, Worker or Consumer
│   │   ├── graph
│   │   ├── grpc
│   │   ├── redis_worker
│   │   └── restful
│   ├── iface              抽象 handler, service, repository
│   ├── model              對應 database 的資料結構
│   │   ├── dto            database 的資料結構
│   │   ├── option         操作 database 各種 查詢, 更新的資料結構
│   ├── service            各種商業邏輯會放在這下面, 後續資料夾可能會依照 domain 做區分
│   └── repository         針對 database, redis 的存取使用進行的分層 
├── platform_gqlgen.yml    產生 graphql 的設定檔文件
└── test                   測試需要用到的文件可能都會擺在這底下
```

## k6 test
system mackbook air 
Apple M1
16 GB
```
     checks.........................: 100.00% ✓ 26637 ✗ 0    
     data_received..................: 3.8 MB  64 kB/s
     data_sent......................: 12 MB   206 kB/s
     http_req_blocked...............: avg=28.23ms  min=58µs   med=212µs    max=7.6s     p(90)=392µs    p(95)=572µs   
     http_req_connecting............: avg=28.19ms  min=46µs   med=178µs    max=7.6s     p(90)=329µs    p(95)=483µs   
     http_req_duration..............: avg=197.35ms min=4.23ms med=154.82ms max=686.8ms  p(90)=426.77ms p(95)=502.88ms
       { expected_response:true }...: avg=197.35ms min=4.23ms med=154.82ms max=686.8ms  p(90)=426.77ms p(95)=502.88ms
     http_req_failed................: 0.00%   ✓ 0     ✗ 26637
     http_req_receiving.............: avg=55.72µs  min=10µs   med=40µs     max=5.94ms   p(90)=86µs     p(95)=115µs   
     http_req_sending...............: avg=54.02µs  min=7µs    med=37µs     max=5.72ms   p(90)=79µs     p(95)=101µs   
     http_req_tls_handshaking.......: avg=0s       min=0s     med=0s       max=0s       p(90)=0s       p(95)=0s      
     http_req_waiting...............: avg=197.24ms min=4.13ms med=154.7ms  max=686.74ms p(90)=426.66ms p(95)=502.77ms
     http_reqs......................: 26637   441.242097/s
     iteration_duration.............: avg=225.73ms min=4.59ms med=155.71ms max=8.05s    p(90)=430.98ms p(95)=512.34ms
     iterations.....................: 26637   441.242097/s
     vus............................: 100     min=100 max=100
     vus_max........................: 100     min=100 max=100
```


Please design and implement a backend system for an online trading platform of Pokémon Trading Card Game.
- [v] This online trading platform trades 4 kinds of cards only: Pikachu, Bulbasaur, Charmander, and Squirtle.
- [v] The price of cards is between 1.00 USD and 10.00 USD.
- [v] Users on this platform are called traders.
- [v] There are 10K traders.
- [v] Traders own unlimited USD and cards.
- [v] Traders can send orders to the platform when they want to buy or sell cards at certain prices.
- [v] A trader can only buy or sell 1 card in 1 order.
- [v] Traders can only buy cards using USD or sell cards for USD.
- [v] Orders are first come first serve.
- [v] There are 2 situations to make a trade:
    - [v] When a buy order is sent to the platform, there exists an uncompleted sell order, whose price is the lowest one among all uncompleted sell orders and less than or equal to the price of the buy order. Then, a trade is made at the price of the sell order. Both buy and sell orders are completed. Otherwise, the buy order is uncompleted.
    - [v] When a sell order is sent to the platform, there exists an uncompleted buy order, whose price is the highest one among all uncompleted buy orders and greater than or equal to the price of the sell order. Then, a trade is made at the price of the buy order. Both buy and sell orders are completed. Otherwise, the sell order is uncompleted.
- [v] Traders can view the status of their latest 50 orders.
- [v] Traders can view the latest 50 trades on each kind of cards.
- [v] If the sequence of orders is fixed, the results must be the same no matter how many times you execute the sequence.
## Basic Requirements:
- [v] RESTful API
- [v] Relational database (PostgreSQL, MySQL, ...)
- [v] Containerize
- [v] Testing
- [v] Gracefully shutdown
## Advanced Requirements:
- [v] Multithreading
- Maximize performance of finishing 1M orders
- OpenAPI (Swagger)
- Set up configs using environment variables
- View logs on visualization dashboard (Kibana, Grafana, ...)
- Microservice
- Message queue (Apache Kafka, Apache Pulsar, ...)
- gRPC
- GraphQL
- [v] Docker Compose
- Kubernetes
- Cloud computing platforms (AWS, Azure, GCP, ...) 
- NoSQL
- CI/CD
- [v] User authentication and authorization
- [v] High availability
- ...

## More
- [v] 可以購買多張卡片 , 以上述的吃單或掛單方式進行撮合

