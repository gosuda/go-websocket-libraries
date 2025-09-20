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
**Last Updated:** Sat, 20 Sep 2025 03:09:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 211.64 | 36.58 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.38 | 27.72 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 437.63 | 17.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 338.05 | 22.96 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 516.38 | 15.06 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)