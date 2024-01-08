# 부록A private-isu 공략 실천

부록A "private-isu 공략 실천"의 샘플 코드입니다.

이 디렉토리의 private-isu/webapp 아래에는 [catatsuy/private-isu](https://github.com/catatsuy/private-isu)를 부록의 설명에 따라 변경한 코드의 예가 저장되어 있습니다.

### unicorn worker 프로세스를 4로 설정(약 13,000점)

- [리스트 A.1 변경한 unicorn_config.rb](private-isu/webapp/ruby/unicorn_config.rb)
- [commit](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/83a57020b2e205c8a7d6163ee3c58fed361f6605)

### 정적 파일을 nginx로 전달(약 17,000점)

- [리스트 A.2 /etc/nginx/sites-available/isucon.conf](private-isu/webapp/etc/nginx/conf.d/default.conf)
- [commit](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/9f34d0bd90146a050cc4c7e13d5c89743c7f77be)

### 업로드 이미지를 정적 파일화(약 22,000점)

- [리스트 A.3 업로드 이미지를 순차적으로 정적 파일로 이전하는 응용 프로그램의 변경](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/907c70b3f7e722068fd3462445d8cee8efb27a76)
- [리스트 A.4 /image/ 이하의 정적 파일이 있으면 전달, 없으면 응용 프로그램 서버에 리버스 프락시하는 nginx 설정(발췌)](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/a49eb2dfd307ffe7a85eb9cfbcae3912e8427f0d)


### posts와 users를 JOIN해 필요한 행 수만 취득(약 90,000점)

- [리스트 A.9, 리스트 A.10의 차이](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/7f73caf78e714e982b9f24478e0919b8e50af2b6)

### 프리페어드 스테이트먼트를 개선(약 110,000점)

- [리스트 A.12 prepare.execute를 xquery로 변경하는 차이점](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/57592ee2681fc3551ab810932b1706fc775aac43)

### posts에서 N+1 쿼리 결과 캐시(약 180,000점)

- [리스트 A.13 make_posts에서 캐시를 다루는 코드의 예](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/8de837130c50186ce5cd08b560552cd97a1e9b34)

### 적절한 인덱스를 사용할 수 없는 쿼리를 해결(약 200,000점)

- [STRAIGHT_JOIN 및 FORCE INDEX를 추가한 예](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/8bf1580d4542dd7e4c798dcc6c6aead6cd0bd339)


### 외부 명령 호출 중지(약 240,000점)

- [openssl 명령 호출을 중지하고 openssl gem 사용](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/8f3b4a8839a3583a8abd2bde0cea2e2bfa2f8c20)

### MySQL 설정 변경 (약 250,000점)

- [MySQL 설정 조정](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/ecabb326243b04fb5e7e669d034ad2eeb6297474)

### memcached에 대한 N + 1 제거 (약 300,000 점)

- [memcached에서 N+1 문제를 get_multi를 사용하여 해결](https://github.com/tatsujin-web-performance/tatsujin-web-performance/commit/ca092dc1e02f7b3b0d79c96b1271bf0dbd4bb5b9)
