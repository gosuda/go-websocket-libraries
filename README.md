# go-websocket-libraries

Go WebSocket Library Comparison

**Libraries Benchmarked:**

- https://github.com/gorilla/websocket
- https://github.com/lxzan/gws
- https://github.com/gobwas/ws
- https://github.com/lesismal/nbio
- https://github.com/coder/websocket

**Latest Benchmark Results:**

<!-- BENCHMARK_TABLE_START -->
**Last Updated:** Fri, 04 Sep 2026 07:25:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 233.62 | 33.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 252.87 | 30.62 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 438.57 | 17.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 308.16 | 25.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 543.64 | 14.30 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)