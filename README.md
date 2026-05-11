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
**Last Updated:** Mon, 11 May 2026 06:24:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 236.36 | 32.66 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.52 | 27.61 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 503.99 | 15.37 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 336.89 | 22.93 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 575.83 | 13.50 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)