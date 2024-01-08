// http 처리의 module을 import
import http from "k6/http";

// HTML을 구문 분석하는 함수를 import
import { parseHTML } from "k6/html";

// 요청 대상 URL을 생성하는 함수를 import
import { url } from "./config.js";

// 파일을 바이너리로 열기
const testImage = open("testimage.jpg", "b");

// k6이 실행하는 함수
// 로그인하여 이미지를 게시하는 시나리오
export default function () {
  const res = http.post(url("/login"), {
    account_name: "terra",
    password: "terraterra",
  });
  const doc = parseHTML(res.body);
  const token = doc.find('input[name="csrf_token"]').first().attr("value");
  http.post(url("/"), {
    // http.file에서 파일 업로드
    file: http.file(testImage, "testimage.jpg", "image/jpeg"),
    body: "Posted by k6",
    csrf_token: token,
  });
}
