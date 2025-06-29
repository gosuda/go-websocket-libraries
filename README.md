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
**Last Updated:** Sun, 29 Jun 2025 03:47:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.31 | 38.24 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.04 | 26.63 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 452.38 | 17.08 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 353.54 | 21.90 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.52 | 14.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)