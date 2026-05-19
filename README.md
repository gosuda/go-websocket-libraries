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
**Last Updated:** Tue, 19 May 2026 06:35:17 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 248.55 | 31.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 303.54 | 25.49 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 524.03 | 14.78 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 376.95 | 20.57 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 614.28 | 12.64 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)