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
**Last Updated:** Thu, 10 Sep 2026 09:51:42 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 253.84 | 30.52 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.63 | 27.55 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 507.09 | 15.25 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 338.04 | 22.87 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 578.81 | 13.40 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)