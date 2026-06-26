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
**Last Updated:** Fri, 26 Jun 2026 06:38:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 283.18 | 27.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 347.16 | 22.30 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 583.17 | 13.28 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 385.83 | 20.09 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 701.44 | 11.08 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)