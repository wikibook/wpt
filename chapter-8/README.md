# Chapter8 알아 두면 좋은 고속화 방법

## 8-1 외부 명령 실행이 아닌 라이브러리 사용

### 리스트 8.1 openssl 명령을 실행하는 Go의 초기 구현

```go
func digest(src string) string {
  out, err := exec.Command("/bin/bash", "-c", `printf "%s" `+escapeshellarg(src)+` | openssl dgst -sha512 | sed 's/^.*= //'`).Output()
  （생략）
```

https://github.com/catatsuy/private-isu/blob/0c9a8f258c759d5133c6200c6453f82703663614/webapp/golang/app.go#L122-L131

### 리스트 8.2 openssl 명령을 실행하는 Ruby의 초기 구현

```ruby
def digest(src)
  `printf "%s" #{Shellwords.shellescape(src)} | openssl dgst -sha512 | sed 's/^.*= //'`.strip
end
```

https://github.com/catatsuy/private-isu/blob/0c9a8f258c759d5133c6200c6453f82703663614/webapp/ruby/app.rb#L78-L81

### 리스트 8.3 Go 구현

```go
import (
  "fmt"
  "crypto/sha512"
（생략）
)

func digest(src string) string {
  return fmt.Sprintf("%x", sha512.Sum512([]byte(src)))
}
```

### 리스트 8.4 Ruby 구현

```ruby
require 'openssl'
（생략）

def digest(src)
  return OpenSSL::Digest::SHA512.hexdigest(src)
end
```

## Column：구현하는 언어에 따라 속도가 빨라지는가?

### 리스트 8.5 strings.newReplacer 사용

```go
r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
fmt.Println(r.Replace("This is <b>HTML</b>!")) // This is &lt;b&gt;HTML&lt;/b&gt;!
```

## 8-2 개발용 설정에서 불필요한 로그를 출력하지 않는다

### 리스트 8.6 디버그 모드 비활성화, 로그 레벨 변경

```diff
 func main() {
  e := echo.New()
- e.Debug = true
- e.Logger.SetLevel(log.DEBUG)
+ e.Debug = false
+ e.Logger.SetLevel(log.ERROR)
```

https://github.com/isucon/isucon11-qualify/blob/1011682c2d5afcc563f4ebf0e4c88a5124f63614/webapp/go/main.go#L211-L212

## 8-3 HTTP 클라이언트 사용 기법

### 리스트 8.7 res.Body.Close()를 실행해 응답의 Body를 읽는다.

```go
res, err := http.DefaultClient.Do(req)
if err != nil {
  log.Fatal(err)
}
defer res.Body.Close()

_, err = io.ReadAll(res.Body)
if err != nil {
  log.Fatal(err)
}
```

### 리스트 8.8 Timeout 지정

``` go
hClient := http.Client{
  Timeout: 5 * time.Second,
}
```

### 리스트 8.9 http.Transportd에서 확인해야 하는 좋은 설정

``` go
hClient := http.Client{
  Timeout:   5 * time.Second,
  Transport: &http.Transport{
    MaxIdleConns:        500,
    MaxIdleConnsPerHost: 200,
    IdleConnTimeout:     120 * time.Second,
  },
}
```

## 8-4 정적 파일을 리버스 프락시에서 직접 전달

### 리스트 8.10 /home/isucon/private_isu/webapp/public/image/ 디렉터리에 이미지 파일 배치

```nginx
server {
  # 省略
  location /image/ {
    root /home/isucon/private_isu/webapp/public/;
    try_files $uri @app;
  }

  location @app {
    proxy_pass http://localhost:8080;
  }
```

## 8-5 클라이언트 측에서 캐시를 활용하기 위해 HTTP 헤더를 사용

### 리스트 8.11 Cache-Control 헤더를 응답에 포함하는 설정

```nginx
server {
  # 省略
  location /image/ {
    root /home/isucon/private_isu/webapp/public/;
    expires 1d;
  }
```
