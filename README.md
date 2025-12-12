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
**Last Updated:** Fri, 12 Dec 2025 03:47:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 220.42 | 35.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.15 | 26.77 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 440.10 | 17.57 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 345.60 | 22.37 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 541.00 | 14.38 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)