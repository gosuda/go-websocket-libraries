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
**Last Updated:** Mon, 12 Jan 2026 04:04:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.26 | 31.80 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.70 | 25.78 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 543.72 | 14.23 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 376.73 | 20.56 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 626.39 | 12.40 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)