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
**Last Updated:** Thu, 17 Sep 2026 07:58:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 250.05 | 30.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 275.96 | 28.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 505.52 | 15.31 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 340.51 | 22.70 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 593.20 | 13.10 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)