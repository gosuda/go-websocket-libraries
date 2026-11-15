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
**Last Updated:** Sun, 15 Nov 2026 09:52:45 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 252.65 | 30.61 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 278.64 | 27.78 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 504.33 | 15.32 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 336.17 | 23.04 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 600.70 | 12.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)