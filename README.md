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
**Last Updated:** Sun, 07 Sep 2025 03:24:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 230.59 | 33.56 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 305.05 | 25.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 467.57 | 16.55 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 364.56 | 21.20 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 563.15 | 13.79 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)