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
**Last Updated:** Sat, 18 Oct 2025 03:17:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 210.43 | 36.69 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.91 | 27.14 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 431.47 | 17.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 334.10 | 23.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 519.73 | 14.96 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)