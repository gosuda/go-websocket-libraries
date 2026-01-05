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
**Last Updated:** Mon, 05 Jan 2026 04:09:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 238.83 | 32.41 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.55 | 26.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 543.95 | 14.24 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 375.46 | 20.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 620.70 | 12.52 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)