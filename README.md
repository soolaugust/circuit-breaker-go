# Circulate Breaker Pattern

[中文](https://mp.weixin.qq.com/s?__biz=Mzg4MTIwMTkxNQ==&mid=2247484553&idx=1&sn=e6c8d380629ecf3b4a8f223518d3c662&chksm=cf68c31df81f4a0b340bf4a4008f93d2db02b0e6eaade0c2cfec2e7b49cc391aa66ace9f41a0&token=753990504&lang=zh_CN#rd)

`Circulate Breaker Pattern` is a modern-day design pattern for microservice. It's used for such situation:

* To prevent an application from trying to invoke a remote service or access a shared resource if this operation is highly likely to fail.