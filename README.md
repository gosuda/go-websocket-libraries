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
**Last Updated:** Sun, 22 Mar 2026 04:34:47 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 290.55 | 26.69 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 403.99 | 19.18 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 638.40 | 12.14 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 461.21 | 16.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.0 | 801.92 | 9.70 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)