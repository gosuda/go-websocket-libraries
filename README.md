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
**Last Updated:** Wed, 19 Aug 2026 03:27:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 334.81 | 23.19 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 369.59 | 21.01 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 633.31 | 12.27 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 448.23 | 17.30 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 761.08 | 10.23 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)