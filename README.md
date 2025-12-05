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
**Last Updated:** Fri, 05 Dec 2025 03:42:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 217.27 | 35.63 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.80 | 27.18 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 440.71 | 17.54 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 338.71 | 22.84 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 542.77 | 14.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)