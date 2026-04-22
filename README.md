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
**Last Updated:** Wed, 22 Apr 2026 05:14:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 226.11 | 34.21 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.18 | 27.91 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 495.95 | 15.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 320.42 | 24.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 584.80 | 13.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)