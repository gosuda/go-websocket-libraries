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
**Last Updated:** Wed, 27 Aug 2025 03:25:24 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.05 | 38.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.75 | 26.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 447.56 | 17.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 341.57 | 22.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 550.50 | 14.13 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)