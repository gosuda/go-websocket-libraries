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
**Last Updated:** Sat, 13 Jun 2026 06:41:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 222.88 | 34.73 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 270.84 | 28.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 486.56 | 15.93 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 320.77 | 24.12 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 544.18 | 14.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)