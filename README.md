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
**Last Updated:** Sat, 03 May 2025 03:26:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 207.74 | 37.20 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.54 | 27.06 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 449.62 | 17.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 356.06 | 21.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 535.62 | 14.48 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)