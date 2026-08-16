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
**Last Updated:** Sun, 16 Aug 2026 03:30:44 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 428.23 | 18.16 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 546.79 | 14.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 872.76 | 8.90 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 608.04 | 12.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 1142.13 | 6.82 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)