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
**Last Updated:** Sat, 22 Nov 2025 03:28:30 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 200.54 | 38.52 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 258.33 | 29.97 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 419.50 | 18.41 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 302.62 | 25.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 496.70 | 15.63 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)