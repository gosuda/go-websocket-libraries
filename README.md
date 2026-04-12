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
**Last Updated:** Sun, 12 Apr 2026 05:13:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 229.56 | 33.69 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.45 | 27.61 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 510.47 | 15.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 336.71 | 22.98 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 581.93 | 13.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)