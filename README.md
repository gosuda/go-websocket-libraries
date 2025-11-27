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
**Last Updated:** Thu, 27 Nov 2025 03:35:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 226.80 | 34.12 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 303.91 | 25.44 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 464.57 | 16.65 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 358.75 | 21.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 552.76 | 14.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)