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
**Last Updated:** Sat, 25 Oct 2025 03:25:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 214.55 | 36.05 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.15 | 26.75 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 460.79 | 16.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 347.71 | 22.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 552.84 | 14.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)