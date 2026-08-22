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
**Last Updated:** Sat, 22 Aug 2026 03:23:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 253.48 | 30.52 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.90 | 27.72 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 514.54 | 15.03 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 341.31 | 22.71 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 596.89 | 13.00 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)