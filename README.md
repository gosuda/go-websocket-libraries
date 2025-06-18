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
**Last Updated:** Wed, 18 Jun 2025 03:38:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.03 | 37.73 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.45 | 26.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 452.79 | 17.04 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 357.67 | 21.62 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 551.43 | 14.09 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)