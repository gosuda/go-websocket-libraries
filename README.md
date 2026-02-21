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
**Last Updated:** Sat, 21 Feb 2026 04:19:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 244.32 | 31.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.99 | 26.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 552.22 | 14.01 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 375.56 | 20.63 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 634.07 | 12.26 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)