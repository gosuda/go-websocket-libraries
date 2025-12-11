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
**Last Updated:** Thu, 11 Dec 2025 03:48:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 220.07 | 35.16 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.99 | 27.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.14 | 17.13 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 347.56 | 22.25 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 549.53 | 14.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)