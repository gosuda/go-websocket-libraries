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
**Last Updated:** Sat, 31 Jan 2026 04:17:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.97 | 35.17 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 273.45 | 28.32 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 498.56 | 15.52 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 338.70 | 22.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 577.55 | 13.46 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)