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
**Last Updated:** Thu, 19 Nov 2026 10:44:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 266.24 | 29.12 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 321.09 | 24.09 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 528.93 | 14.66 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 365.71 | 21.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 652.92 | 11.91 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)