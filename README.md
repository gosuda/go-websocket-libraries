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
**Last Updated:** Sun, 22 Jun 2025 03:43:58 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.08 | 37.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.63 | 26.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 460.97 | 16.76 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 359.06 | 21.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 554.67 | 13.99 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)