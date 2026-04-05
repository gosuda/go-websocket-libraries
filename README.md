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
**Last Updated:** Sun, 05 Apr 2026 05:02:45 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 245.17 | 31.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.53 | 26.52 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 521.19 | 14.86 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 357.58 | 21.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 601.09 | 12.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)