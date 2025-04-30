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
**Last Updated:** Wed, 30 Apr 2025 02:33:35 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 210.95 | 36.67 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 306.39 | 25.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 463.95 | 16.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 363.68 | 21.29 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 548.96 | 14.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)