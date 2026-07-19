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
**Last Updated:** Sun, 19 Jul 2026 05:38:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 252.62 | 30.63 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.34 | 27.97 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 494.63 | 15.65 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 336.92 | 22.99 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 573.49 | 13.55 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)