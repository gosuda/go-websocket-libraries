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
**Last Updated:** Sat, 06 Sep 2025 03:08:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 199.76 | 38.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.40 | 27.00 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 457.01 | 16.93 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 354.03 | 21.86 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 558.71 | 13.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)