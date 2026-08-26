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
**Last Updated:** Wed, 26 Aug 2026 03:36:47 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 252.78 | 30.61 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 263.77 | 29.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 494.06 | 15.68 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 339.95 | 22.77 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 573.49 | 13.54 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)