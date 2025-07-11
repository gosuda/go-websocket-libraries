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
**Last Updated:** Fri, 11 Jul 2025 03:49:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.37 | 38.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.50 | 26.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 450.64 | 17.16 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 349.66 | 22.12 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 544.72 | 14.28 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)