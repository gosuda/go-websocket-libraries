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
**Last Updated:** Wed, 22 Oct 2025 03:35:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 220.85 | 35.07 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.21 | 25.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.57 | 17.06 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 348.95 | 22.19 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 542.50 | 14.33 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)