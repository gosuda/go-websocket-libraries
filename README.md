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
**Last Updated:** Mon, 23 Jun 2025 03:45:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 203.89 | 37.90 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.23 | 26.32 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 452.04 | 17.09 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 348.24 | 22.21 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 548.20 | 14.17 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)