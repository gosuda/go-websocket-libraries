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
**Last Updated:** Sat, 27 Jun 2026 06:10:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 259.38 | 29.86 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.06 | 26.66 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 507.71 | 15.25 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 342.02 | 22.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 571.53 | 13.58 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)