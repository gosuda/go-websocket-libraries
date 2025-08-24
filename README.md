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
**Last Updated:** Sun, 24 Aug 2025 03:38:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 200.27 | 38.61 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.10 | 26.81 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 470.13 | 16.43 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 344.39 | 22.49 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 553.63 | 14.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)