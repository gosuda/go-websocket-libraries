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
**Last Updated:** Wed, 12 Nov 2025 03:35:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 222.66 | 34.78 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.12 | 26.13 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 457.95 | 16.88 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 352.12 | 21.96 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 556.21 | 13.99 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)