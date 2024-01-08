# Chapter4 시나리오 부하 테스트

Chapter4 "시나리오 부하 테스트"의 샘플 코드입니다.

## 4-2 k6로 간단한 부하 테스트

- 리스트 4.1 단일 URL로 요청을 전송하는 시나리오 [`ab.js`](ab.js)

## 4-3 k6로 시나리오 작성

- 리스트 4.3 대상 URL을 지정하는 함수를 정의한 config.js [`config.js`](config.js)
- 리스트 4.4 웹 서비스를 초기화하는 시나리오 initialize.js [`initialize.js`](initialize.js)
- 리스트 4.5 사용자가 로그인해 댓글을 게시하는 시나리오 comment.js [`comment.js`](comment.js)
- 리스트 4.7 로그인한 후 양식에서 이미지를 업로드하는 시나리오 postimage.js [`postimage.js`](postimage.js)
- 리스트 4.8 계정 정보를 정의하는 JSON 파일 accounts.json [`accounts.json`](accounts.json)
- 리스트 4.9 accounts.json을 SharedArray로 로드하는 모듈 accounts.js [`accounts.js`](accounts.js)

(주) 리스트 4.9 [`accounts.js`] (accounts.js)를`import`로`getAccount()`함수를 사용하는 시나리오에서 변경 사항을 [`comment.js`](comment.js ) 와 [`postimage.js`](postimage.js) 에 대해서 적용하는 경우는, 이 디렉토리에서 이하의 커멘드를 실행해 patch 를 적용해 주세요.

```console
$ patch -p2 < comment.js.patch
$ patch -p2 < postimage.js.patch
```

## 4-4 여러 시나리오를 결합한 통합 시나리오 실시

- 리스트 4.11 여러 시나리오 함수를 조합해 실행 integrated.js [`integrated.js`](integrated.js)

(주) [`integrated.js`](integrated.js) 는 [`comments.js`](comments.js) 와 [`postimage.js`](postimage.js) 에 상기 patch가 적용된 것을 전제로 한다.
