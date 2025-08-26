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
**Last Updated:** Tue, 26 Aug 2025 03:32:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.22 | 38.66 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.44 | 26.49 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 464.89 | 16.63 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 359.59 | 21.53 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 553.93 | 14.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)