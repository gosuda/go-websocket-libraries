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
**Last Updated:** Fri, 12 Sep 2025 03:09:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 195.20 | 39.57 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 258.98 | 29.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 423.49 | 18.26 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 305.38 | 25.30 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 507.16 | 15.31 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)