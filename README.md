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
**Last Updated:** Thu, 15 Oct 2026 07:45:16 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 239.10 | 32.35 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 268.38 | 28.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 483.08 | 16.05 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 326.44 | 23.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 546.69 | 14.17 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)