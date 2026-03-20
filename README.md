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
**Last Updated:** Fri, 20 Mar 2026 04:28:00 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 233.82 | 33.13 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.08 | 26.51 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 512.43 | 15.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 360.37 | 21.55 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.0 | 585.23 | 13.27 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)