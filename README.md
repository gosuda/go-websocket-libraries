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
**Last Updated:** Tue, 03 Feb 2026 04:28:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 232.82 | 33.24 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.33 | 27.74 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 524.64 | 14.76 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 350.10 | 22.07 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 603.90 | 12.89 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)