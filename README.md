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
**Last Updated:** Wed, 25 Feb 2026 04:34:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 228.12 | 33.93 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.40 | 27.73 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 497.64 | 15.59 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 341.79 | 22.66 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 571.30 | 13.59 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)