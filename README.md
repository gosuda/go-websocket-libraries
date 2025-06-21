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
**Last Updated:** Sat, 21 Jun 2025 03:35:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.26 | 38.27 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.24 | 26.31 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 448.60 | 17.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 341.03 | 22.68 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 542.59 | 14.32 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)