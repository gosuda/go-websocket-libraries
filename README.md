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
**Last Updated:** Sat, 15 Aug 2026 03:21:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 251.33 | 30.80 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.04 | 27.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 510.07 | 15.18 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 332.82 | 23.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 583.38 | 13.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)