# Chapter9 OS 기초 지식과 튜닝

## 9-8 리눅스 커널 매개변수

### 리스트 9.1 UNIX domain socket을 사용한 예

```nginx
server {

  ## 80번 포트에서 접속 대기할 때의 설정(#을 붙여 주석 처리)
  # listen 80;

  ## /var/run/nginx.sock で接続を待機する際の設定
  listen unix:/var/run/nginx.sock;

<이하 생략>
```

### 리스트 9.2 unicorn_config.rb의 설정

```ruby
worker_processes 1
preload_app true
listen "0.0.0.0:8080"
```

### 리스트 9.3 /tmp/webapp.sock으로 변경

```ruby
worker_processes 1
preload_app true
listen "/tmp/webapp.sock"
```

### 리스트 9.4 Go 구현 재작성

```go
## "/tmp/webapp.sock" で listen(2) する
listener, err := net.Listen("unix", "/tmp/webapp.sock")
if err != nil {
        log.Fatalf("Failed to listen on /tmp/webapp.sock: %s.", err)
}
defer func() {
        err := listener.Close()
        if err != nil {
                log.Fatalf("Failed to close listener: %s.", err)
        }
}()

## systemd 등으로부터 송신되는 시그널을 받는다
c := make(chan os.Signal, 2)
signal.Notify(c, os.Interrupt, syscall.SIGTERM)
go func() {
        <-c
        err := listener.Close()
        if err != nil {
                log.Fatalf("Failed to close listener: %s.", err)
        }
}()

log.Fatal(http.Serve(listener, mux))
```

### 리스트 9.5 초기 상태 설정

```nginx
server {
<省略>
  location / {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://localhost:8080;
  }
}  
```

### 리스트 9.6 /tmp/webapp.sock을 업스트림 서버로 지정한 설정

```nginx
upstream webapp {
  server unix:/tmp/webapp.sock;
}

server {
<생략>
  location / {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://webapp;
  }
}
```
