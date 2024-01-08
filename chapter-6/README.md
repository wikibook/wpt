# Chapter6 리버스 프락시 사용

Chapter5 「데이터베이스 튜닝」의 샘플 코드입니다.

### 6-3 nginx란?

### 리스트 6.1 /etc/nginx/nginx.conf 설정

```nginx
include /etc/nginx/conf.d/*.conf;
include /etc/nginx/sites-enabled/*;
```

### 리스트 6.2 /etc/nginx/sites-available/isucon.conf 설정

```nginx
server {
  listen 80;

  client_max_body_size 10m;
  root /home/isucon/private_isu/webapp/public/;

  location / {
    proxy_set_header Host $host;
    proxy_pass http://localhost:8080;
  }
}
```

https://github.com/catatsuy/private-isu/blob/master/provisioning/image/files/etc/nginx/sites-available/isucon.conf

### 리스트 6.3 nginx를 통해 정적 파일 배포

```nginx
server {
  listen 80;

  # 省略

  location /css/ {
    root /home/isucon/private_isu/webapp/public/;
  }

  location /js/ {
    root /home/isucon/private_isu/webapp/public/;
  }

  location / {
    proxy_set_header Host $host;
    proxy_pass http://localhost:8080;
  }
}
```

### 리스트 6.4 expires 설정

```nginx
  location /css/ {
    root /home/isucon/private_isu/webapp/public/;
    expires 1d;
  }
```

## 6-5 nginx로 전송할 때 데이터 압축

### 리스트 6.5 gzip 압축을 사용하는 경우의 설정

```nginx
gzip on;
gzip_types text/css text/javascript application/javascript application/x-javascript application/json;
gzip_min_length 1k;
```

## 6-7 nginx와 업스트림 서버의 연결 관리

### 리스트 6.6 업스트림 서버와의 연결을 유지하는 설정

```nginx
location / {
  proxy_http_version 1.1;
  proxy_set_header Connection "";
  proxy_pass http://app;
}
```

### 리스트 6.7 keepalive 및 keepalive_requests 사용

```nginx
upstream app {
  server localhost:8080;

  keepalive 32;
  keepalive_requests 10000;
}
```

## Column : 더 빨라진 nginx의 속도

### 리스트 6.8 sendfile과 tcp_nopush를 모두 활성화하는 설정

```nginx
sendfile on;
tcp_nopush on;
```
