# SLx API

SLA & SRE 를 위한 API Project

# SLA(Service Level Agreement)
사전에 정의된 수준의 서비스를 제공하기로 맺는 협약

```shell
Downtime : 기간(daily, monthly..) 중에 서비스를 사용할 수 없는 시간
```
```shell
Uptime : 기간(daily, monthly..) 중에 서비스를 사용할 수 있는 시간
```

서비스 제공업체는 다음 예와 같이 특정 기간에 대한 가동율(퍼센티지)를 보장한다는 것과 같은 내용을 게시

"월별 가동율 99.9%", 이 말은 제공하는 서비스가 월별 가동 시간에 99.9%만큼 서비스를 사용할 수 있게끔 제공한다는 것을 의미
```shell
월별 가동율 99.9%가 허용하는 downtime은 43분 49초이다
```

기간별 가동율 99.9%가 허용하는 downtime
```shell
"daily": "1m26s"
"weekly": "10m4s"
"monthly": "43m49s"
"yearly": "8h45m56s"
```

## SLA API
퍼센티지를 파라미터로 받아 가동율에 대한 uptime, downtime을 반환한다
```shell
http://localhost:8000/sla/99.9

{
  "uptime": {
    "daily": "23h58m33s",
    "weekly": "6d 23h49m55s",
    "monthly": "30d 9h45m16s",
    "yearly": "364d 21h3m15s"
  },
  "downtime": {
    "daily": "1m26s",
    "weekly": "10m4s",
    "monthly": "43m49s",
    "yearly": "8h45m56s"
  }
}
```

### Run
No build
```shell
$ go run main.go
```

build
```shell
$ go build main.go
$ ./main
```

### Docker
Image Build
```shell
$ docker build -t [CONTAINER_REPOSITORY]:[CONTAINER_VERSION] .
```

Container Run
```shell
$ docker run -d --name [CONTAINER_NAME] -p 8000:8000 --restart=always [CONTAINER_REPOSITORY]:[CONTAINER_VERSION]
```
