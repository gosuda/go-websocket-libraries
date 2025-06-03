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
**Last Updated:** Tue, 03 Jun 2025 03:37:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 192.65 | 40.10 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 275.44 | 28.04 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 426.70 | 18.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 332.81 | 23.21 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 507.47 | 15.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)