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
**Last Updated:** Thu, 30 Jul 2026 05:25:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 339.36 | 22.87 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 375.70 | 20.62 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 640.85 | 12.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 455.87 | 17.04 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 753.48 | 10.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)