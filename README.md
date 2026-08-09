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
**Last Updated:** Sun, 09 Aug 2026 04:03:47 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 452.22 | 17.19 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 615.95 | 12.62 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 898.24 | 8.64 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 654.51 | 11.87 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 1215.79 | 6.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)