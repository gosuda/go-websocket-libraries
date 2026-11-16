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
**Last Updated:** Mon, 16 Nov 2026 12:07:17 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 244.39 | 31.62 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 266.98 | 28.91 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 481.03 | 16.09 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 326.97 | 23.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 560.62 | 13.85 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)