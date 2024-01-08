// SharedArray 를import
import { SharedArray } from "k6/data";

// accounts.json을 로드하여 SharedArray로 설정
const accounts = new SharedArray("accounts", function () {
  return JSON.parse(open("./accounts.json"));
});

// SharedArray에서 무작위로 1개를 가져와 반환하는 함수
export function getAccount() {
  return accounts[Math.floor(Math.random() * accounts.length)];
}
