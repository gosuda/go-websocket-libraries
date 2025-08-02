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
**Last Updated:** Sat, 02 Aug 2025 03:46:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.80 | 37.43 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 307.17 | 25.19 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 458.41 | 16.86 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 355.01 | 21.81 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 547.08 | 14.21 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)