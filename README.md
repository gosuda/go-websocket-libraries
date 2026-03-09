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
**Last Updated:** Mon, 09 Mar 2026 04:32:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 222.25 | 34.73 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 266.25 | 29.08 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 473.75 | 16.33 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 342.61 | 22.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 556.09 | 13.97 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)