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
**Last Updated:** Sat, 07 Jun 2025 03:35:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 204.64 | 37.72 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.62 | 26.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 467.70 | 16.51 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 344.71 | 22.42 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 552.80 | 14.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)