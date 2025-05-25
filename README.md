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
**Last Updated:** Sun, 25 May 2025 03:38:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 196.07 | 39.45 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.59 | 26.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 450.72 | 17.15 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 352.59 | 21.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 548.67 | 14.17 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)