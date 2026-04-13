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
**Last Updated:** Mon, 13 Apr 2026 05:30:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 221.25 | 34.89 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 281.70 | 27.44 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 489.65 | 15.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 335.15 | 23.09 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 578.30 | 13.46 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)