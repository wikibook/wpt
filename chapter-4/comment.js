// k6m에서 http처리 module을 import
import http from "k6/http";

// k6에서 check 함수를 import
import { check } from "k6";

// k6에서 HTML을 구문 분석하는 함수를 import
import { parseHTML } from "k6/html";

// url 함수를 import
import { url } from "./config.js";

// 벤치마커가 실시하는 시나리오 함수
// 로그인 후 댓글을 게시
export default function () {
  // /login에 대해 계정 이름과 암호 전송
  const login_res = http.post(url("/login"), {
    account_name: "terra",
    password: "terraterra",
  });

  // 응답 상태 코드가 200인지 확인
  check(login_res, {
    "is status 200": (r) => r.status === 200,
  });

  // 사용자 페이지 /@terra 를 GET
  const res = http.get(url("/@terra"));

  // 응답 내용을 HTML로 해석
  const doc = parseHTML(res.body);

  // 양식의 hidden 요소에서 csrf_token, post_id 추출
  const token = doc.find('input[name="csrf_token"]').first().attr("value");
  const post_id = doc.find('input[name="post_id"]').first().attr("value");

  // /comment에 대해 post_id, csrf_token과 함께 주석 본문을 POST
  const comment_res = http.post(url("/comment"), {
    post_id: post_id,
    csrf_token: token,
    comment: "Hello k6!",
  });
  check(comment_res, {
    "is status 200": (r) => r.status === 200,
  });
}
