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
**Last Updated:** Sat, 31 May 2025 03:32:11 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 191.39 | 40.31 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.15 | 27.77 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 421.40 | 18.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 334.04 | 23.13 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 510.69 | 15.19 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)