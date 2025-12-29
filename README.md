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
**Last Updated:** Mon, 29 Dec 2025 04:03:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 244.31 | 31.66 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 302.07 | 25.68 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 543.79 | 14.23 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 384.17 | 20.15 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 629.55 | 12.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)