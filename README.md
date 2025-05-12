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
**Last Updated:** Mon, 12 May 2025 03:34:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.54 | 38.47 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.54 | 25.71 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 461.70 | 16.73 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 355.83 | 21.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 557.64 | 13.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)