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
**Last Updated:** Wed, 30 Apr 2025 02:35:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.79 | 37.30 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.13 | 26.68 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 455.03 | 16.96 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 359.08 | 21.56 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 540.91 | 14.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)