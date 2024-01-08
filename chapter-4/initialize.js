// k6의http처리module을import
import http from "k6/http";

// k6의sleep함수를import
import { sleep } from "k6";

// 독자적으로 정의한 url 함수를 import
import { url } from "./config.js";

// k6이 실행하는 함수
// /initialize 에 10초의 타임아웃을 지정해 GET요청해, 완료 후 1초 대기한다
export default function () {
  http.get(url("/initialize"), {
    timeout: "10s",
  });
  sleep(1);
}
