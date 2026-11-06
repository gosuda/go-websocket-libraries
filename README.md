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
**Last Updated:** Fri, 06 Nov 2026 09:49:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 323.85 | 24.01 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 336.37 | 23.06 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 617.88 | 12.57 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 453.09 | 17.13 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 708.91 | 11.00 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)