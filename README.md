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
**Last Updated:** Thu, 07 Aug 2025 03:57:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 190.26 | 40.66 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 281.15 | 27.45 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 414.57 | 18.65 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 322.52 | 23.98 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 490.45 | 15.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)