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
**Last Updated:** Sun, 10 May 2026 05:52:10 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 234.61 | 33.02 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.05 | 27.16 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 513.55 | 15.07 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 339.28 | 22.84 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 577.50 | 13.45 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)