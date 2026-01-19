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
**Last Updated:** Mon, 19 Jan 2026 04:05:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 241.51 | 32.04 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 301.08 | 25.73 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 530.95 | 14.58 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 376.33 | 20.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 614.64 | 12.65 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)