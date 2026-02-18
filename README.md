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
**Last Updated:** Wed, 18 Feb 2026 04:34:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 290.38 | 26.71 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 395.72 | 19.59 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 644.00 | 12.04 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 452.97 | 17.13 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 799.43 | 9.73 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)