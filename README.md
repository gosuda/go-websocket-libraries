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
**Last Updated:** Mon, 16 Feb 2026 04:39:34 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 229.63 | 33.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.41 | 27.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 502.86 | 15.39 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 351.97 | 22.01 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 564.68 | 13.78 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)