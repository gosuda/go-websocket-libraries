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
**Last Updated:** Mon, 25 May 2026 06:54:26 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 248.95 | 31.15 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 275.60 | 28.11 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 513.47 | 15.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 375.88 | 20.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 592.59 | 13.14 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)