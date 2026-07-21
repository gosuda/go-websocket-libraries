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
**Last Updated:** Tue, 21 Jul 2026 05:37:13 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 260.11 | 29.74 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.42 | 27.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 481.07 | 16.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 353.82 | 21.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 583.37 | 13.32 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)