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
**Last Updated:** Thu, 02 Apr 2026 04:54:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 298.87 | 25.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 353.92 | 21.93 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 606.15 | 12.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 418.37 | 18.54 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 696.04 | 11.18 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)