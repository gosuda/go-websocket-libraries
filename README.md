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
**Last Updated:** Tue, 10 Nov 2026 10:20:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 241.61 | 32.03 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 256.31 | 30.06 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 465.67 | 16.59 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 315.56 | 24.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 538.98 | 14.42 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)