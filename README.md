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
**Last Updated:** Thu, 04 Jun 2026 07:09:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 248.12 | 31.20 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.42 | 26.27 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 506.04 | 15.31 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 362.27 | 21.38 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 601.78 | 12.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)