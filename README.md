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
**Last Updated:** Wed, 01 Oct 2025 03:33:42 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 227.09 | 34.13 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 332.30 | 23.33 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 467.02 | 16.59 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 386.85 | 20.03 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 612.33 | 12.72 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)