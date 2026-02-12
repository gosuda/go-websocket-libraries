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
**Last Updated:** Thu, 12 Feb 2026 04:37:44 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 244.18 | 31.75 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 303.01 | 25.58 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 528.92 | 14.63 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 373.40 | 20.72 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 631.17 | 12.32 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)