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
**Last Updated:** Wed, 16 Jul 2025 03:48:47 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 204.12 | 37.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.62 | 26.17 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.79 | 17.04 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 349.44 | 22.19 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 538.74 | 14.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)