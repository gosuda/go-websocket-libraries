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
**Last Updated:** Thu, 12 Mar 2026 04:27:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 239.10 | 32.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.00 | 27.17 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 509.66 | 15.20 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 360.09 | 21.51 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 593.91 | 13.10 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)