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
**Last Updated:** Mon, 20 Apr 2026 05:28:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 248.10 | 31.20 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.25 | 26.56 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 534.80 | 14.49 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 364.71 | 21.23 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 621.94 | 12.49 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)