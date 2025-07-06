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
**Last Updated:** Sun, 06 Jul 2025 03:45:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 207.13 | 37.34 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.79 | 25.70 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 472.47 | 16.35 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 364.05 | 21.25 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 574.59 | 13.52 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)