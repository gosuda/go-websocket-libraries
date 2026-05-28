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
**Last Updated:** Thu, 28 May 2026 06:39:16 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 241.17 | 31.86 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 302.96 | 25.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 521.95 | 14.84 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 362.01 | 21.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 606.62 | 12.81 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)