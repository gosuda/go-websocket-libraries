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
**Last Updated:** Sun, 22 Feb 2026 04:31:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 230.47 | 33.53 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.81 | 27.19 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 511.94 | 15.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 346.17 | 22.35 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 579.67 | 13.40 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)