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
**Last Updated:** Thu, 24 Jul 2025 03:50:52 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 207.64 | 37.17 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 304.91 | 25.34 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 461.42 | 16.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 360.88 | 21.46 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 551.58 | 14.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)