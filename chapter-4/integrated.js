// 각 파일에서 시나리오 함수 import
import initialize from "./initialize.js";
import comment from "./comment.js";
import postimage from "./postimage.js";

// k6이 각 함수를 실행할 수 있도록 export
export { initialize, comment, postimage };

// 여러 시나리오를 결합하여 실행하는 옵션 정의
export const options = {
  scenarios: {
    initialize: {
      executor: "shared-iterations", // 일정량의 실행을 여러 VUs에서 공유하는 실행 메커니즘
      vus: 1, // 동시 실행수(초기화이므로 1)
      iterations: 1, // 반복횟수(초기화이므로1회만)
      exec: "initialize", // 실행할 시나리오의 함수명
      maxDuration: "10s", // 최대 실행 시간
    },
    comment: {
      executor: "constant-vus", // 여러 VUs를 병렬로 이동하는 실행 메커니즘
      vus: 4, // 4 VUs에서 실행
      duration: "30s", // 30초 동안 실행
      exec: "comment", // comment 함수 실행
      startTime: "12s", // 12초 후 실행 시작
    },
    postImage: {
      executor: "constant-vus",
      vus: 2,
      duration: "30s",
      exec: "postimage",
      startTime: "12s",
    },
  },
};

// k6이 실행하는 함수. 정의는 비어 있어도 된다
export default function () {}
