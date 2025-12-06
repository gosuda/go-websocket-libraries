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
**Last Updated:** Sat, 06 Dec 2025 03:34:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 222.51 | 34.71 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.04 | 27.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 463.24 | 16.69 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 351.89 | 22.01 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 553.09 | 14.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)