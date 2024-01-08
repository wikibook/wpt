#!/bin/sh

# 실행 시점의 날짜와 시간을 YYYYMMDD-HHMMSS 형식으로 추가한 파일 이름으로 로테이션한다
mv /var/log/nginx/access.log /var/log/nginx/access.log.`date +%Y%m%d-%H%M%S`

# nginx 로그 파일을 다시 열기 위한 신호를 보낸다
nginx -s reopen
