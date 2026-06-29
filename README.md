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
**Last Updated:** Mon, 29 Jun 2026 07:14:47 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 256.08 | 30.21 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.03 | 27.27 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 511.05 | 15.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 343.92 | 22.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 574.83 | 13.51 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)