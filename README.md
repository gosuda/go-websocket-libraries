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
**Last Updated:** Wed, 04 Mar 2026 04:22:56 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 227.54 | 34.03 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 268.43 | 28.85 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 499.23 | 15.50 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 343.40 | 22.54 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 568.12 | 13.69 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)