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
**Last Updated:** Sun, 01 Mar 2026 04:36:10 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 238.49 | 32.46 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.13 | 27.03 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 522.31 | 14.82 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 364.74 | 21.24 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 610.01 | 12.74 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)