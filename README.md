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
**Last Updated:** Mon, 09 Jun 2025 03:42:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.78 | 38.17 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.34 | 26.11 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 445.98 | 17.32 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 354.36 | 21.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 538.04 | 14.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)