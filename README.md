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
**Last Updated:** Mon, 29 Sep 2025 03:28:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.21 | 35.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.83 | 26.70 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.54 | 17.16 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 327.79 | 23.62 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 553.46 | 14.04 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)