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
**Last Updated:** Fri, 18 Jul 2025 03:50:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.66 | 38.31 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 307.32 | 25.16 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 449.74 | 17.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 356.50 | 21.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 548.56 | 14.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)