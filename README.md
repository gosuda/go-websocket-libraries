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
**Last Updated:** Mon, 19 May 2025 03:39:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.51 | 38.31 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.02 | 26.62 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 447.57 | 17.24 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 342.66 | 22.56 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 537.39 | 14.44 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)