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
**Last Updated:** Tue, 19 Aug 2025 03:32:58 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 198.19 | 38.94 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.48 | 27.02 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.84 | 17.07 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 356.26 | 21.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 551.86 | 14.09 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)