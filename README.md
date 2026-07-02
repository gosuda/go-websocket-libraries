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
**Last Updated:** Thu, 02 Jul 2026 06:25:25 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 257.99 | 29.99 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.47 | 27.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 507.72 | 15.25 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 341.07 | 22.68 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 578.22 | 13.45 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)