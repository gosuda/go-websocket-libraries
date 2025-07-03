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
**Last Updated:** Thu, 03 Jul 2025 03:42:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.57 | 37.62 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.12 | 25.84 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 462.04 | 16.70 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 361.62 | 21.43 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 553.65 | 14.06 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)