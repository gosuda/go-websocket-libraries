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
**Last Updated:** Thu, 08 May 2025 03:32:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.06 | 37.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.31 | 26.25 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 456.15 | 16.94 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 361.18 | 21.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 555.88 | 13.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)