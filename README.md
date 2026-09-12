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
**Last Updated:** Sat, 12 Sep 2026 07:21:11 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 246.22 | 31.48 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 273.31 | 28.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 490.66 | 15.77 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 340.26 | 22.79 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 575.04 | 13.50 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)