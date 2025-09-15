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
**Last Updated:** Mon, 15 Sep 2025 03:30:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 205.09 | 37.67 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 262.42 | 29.48 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 412.33 | 18.76 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 320.44 | 24.12 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 491.61 | 15.82 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)