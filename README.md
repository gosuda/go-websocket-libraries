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
**Last Updated:** Sun, 01 Jun 2025 03:51:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 198.64 | 38.91 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.33 | 27.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 436.67 | 17.69 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 339.00 | 22.81 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 532.84 | 14.52 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)