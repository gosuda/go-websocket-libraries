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
**Last Updated:** Thu, 29 Oct 2026 08:53:26 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 250.52 | 30.91 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.01 | 26.57 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 473.13 | 16.36 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 324.17 | 23.87 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 591.00 | 13.13 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)