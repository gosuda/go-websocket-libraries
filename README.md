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
**Last Updated:** Thu, 23 Apr 2026 05:18:02 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.77 | 31.74 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.30 | 26.49 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 517.73 | 14.98 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 363.15 | 21.30 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 594.69 | 13.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)