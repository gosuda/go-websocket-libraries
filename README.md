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
**Last Updated:** Sun, 13 Sep 2026 07:39:40 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 262.20 | 29.52 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.57 | 27.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 518.38 | 14.95 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 348.61 | 22.20 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 598.30 | 12.98 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)