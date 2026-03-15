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
**Last Updated:** Sun, 15 Mar 2026 04:56:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 233.74 | 33.07 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.33 | 27.15 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 498.03 | 15.54 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 352.14 | 21.98 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 568.14 | 13.69 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)