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
**Last Updated:** Mon, 22 Dec 2025 03:57:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 235.13 | 32.94 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.70 | 26.93 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 539.40 | 14.34 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 351.83 | 21.99 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 609.55 | 12.74 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)