// localhost 이외를 대상으로 하는 경우는 여기를 수정
const BASE_URL = "http://localhost";

export function url(path) {
  return `${BASE_URL}${path}`;
}
