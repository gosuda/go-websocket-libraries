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
**Last Updated:** Wed, 25 Mar 2026 04:34:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 255.84 | 30.29 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 308.71 | 25.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 557.00 | 13.90 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 391.23 | 19.79 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.0 | 654.89 | 11.86 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)