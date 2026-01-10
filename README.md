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
**Last Updated:** Sat, 10 Jan 2026 03:48:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 228.07 | 33.94 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.22 | 27.82 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 509.88 | 15.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 345.15 | 22.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 578.02 | 13.46 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)