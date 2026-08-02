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
**Last Updated:** Sun, 02 Aug 2026 05:43:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 274.39 | 28.22 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.05 | 26.21 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 531.63 | 14.59 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 374.63 | 20.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 618.79 | 12.58 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)