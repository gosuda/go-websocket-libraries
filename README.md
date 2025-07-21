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
**Last Updated:** Mon, 21 Jul 2025 03:56:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 197.27 | 39.21 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.90 | 27.53 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 440.11 | 17.50 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 334.61 | 23.12 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 526.53 | 14.75 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)