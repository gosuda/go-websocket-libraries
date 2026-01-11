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
**Last Updated:** Sun, 11 Jan 2026 04:04:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 221.95 | 34.87 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.07 | 27.94 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 503.51 | 15.37 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 335.16 | 23.09 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 573.66 | 13.54 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)