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
**Last Updated:** Fri, 29 Aug 2025 03:25:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.15 | 38.30 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.93 | 26.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.88 | 17.03 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 349.68 | 22.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 545.29 | 14.26 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)