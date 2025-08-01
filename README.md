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
**Last Updated:** Fri, 01 Aug 2025 04:07:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.90 | 38.17 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.54 | 26.29 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 441.01 | 17.52 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 342.53 | 22.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.41 | 14.37 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)