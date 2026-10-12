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
**Last Updated:** Mon, 12 Oct 2026 09:43:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 342.20 | 22.71 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 352.31 | 22.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 631.93 | 12.27 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 442.96 | 17.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 791.09 | 9.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)