# Chapter2 모니터링

## 2-7 모니터링 실시

### 리스트2.1 CPU 사용 시간을 표시하는 쿼리

```
avg without(cpu) (rate(node_cpu_seconds_total{mode!="idle"}[1m]))
```
