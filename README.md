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
**Last Updated:** Tue, 20 Jan 2026 03:59:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 233.62 | 33.12 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.83 | 26.89 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 532.51 | 14.53 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 361.74 | 21.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 613.85 | 12.66 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)