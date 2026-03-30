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
**Last Updated:** Mon, 30 Mar 2026 05:13:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 247.45 | 31.27 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 306.82 | 25.13 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 542.79 | 14.27 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 372.09 | 20.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 629.52 | 12.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)