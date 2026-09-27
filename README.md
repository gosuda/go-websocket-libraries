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
**Last Updated:** Sun, 27 Sep 2026 08:33:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 263.33 | 29.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.75 | 27.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 532.43 | 14.55 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.0 | 349.36 | 22.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 620.30 | 12.54 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)