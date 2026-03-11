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
**Last Updated:** Wed, 11 Mar 2026 04:23:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 232.90 | 33.17 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.52 | 27.34 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.01 | 15.02 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 355.84 | 21.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 582.17 | 13.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)