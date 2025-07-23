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
**Last Updated:** Wed, 23 Jul 2025 03:51:52 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.24 | 37.66 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.90 | 26.13 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 463.46 | 16.67 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 353.30 | 21.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 543.49 | 14.31 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)