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
**Last Updated:** Sun, 14 Jun 2026 07:09:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 250.14 | 30.94 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 304.10 | 25.45 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 507.47 | 15.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 373.12 | 20.76 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 613.25 | 12.69 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)