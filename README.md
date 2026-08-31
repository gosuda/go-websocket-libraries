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
**Last Updated:** Mon, 31 Aug 2026 08:54:08 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 391.30 | 19.85 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 485.31 | 15.99 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 782.71 | 9.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 565.88 | 13.72 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 989.79 | 7.86 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)