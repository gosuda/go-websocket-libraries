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
**Last Updated:** Sat, 07 Mar 2026 04:13:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.34 | 34.25 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 260.51 | 29.70 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 452.45 | 17.07 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 329.63 | 23.48 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 560.46 | 13.88 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)