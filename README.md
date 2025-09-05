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
**Last Updated:** Fri, 05 Sep 2025 03:20:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 197.21 | 39.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.83 | 27.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 446.43 | 17.32 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 343.64 | 22.53 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 534.84 | 14.54 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)