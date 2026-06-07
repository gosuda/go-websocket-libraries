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
**Last Updated:** Sun, 07 Jun 2026 06:50:45 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 241.62 | 32.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.90 | 26.58 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.62 | 15.03 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 357.56 | 21.63 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 612.45 | 12.71 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)