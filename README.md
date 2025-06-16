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
**Last Updated:** Mon, 16 Jun 2025 03:42:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 198.81 | 38.83 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.62 | 26.84 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 448.45 | 17.26 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 344.66 | 22.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 532.00 | 14.59 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)