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
**Last Updated:** Sun, 18 Oct 2026 07:25:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 333.51 | 23.27 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 417.08 | 18.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 695.74 | 11.14 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 474.16 | 16.36 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 899.09 | 8.64 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)