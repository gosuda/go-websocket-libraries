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
**Last Updated:** Thu, 11 Sep 2025 03:22:16 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 215.90 | 35.83 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.46 | 27.61 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 443.02 | 17.46 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 342.15 | 22.64 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 515.97 | 15.06 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)