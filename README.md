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
**Last Updated:** Mon, 14 Jul 2025 03:54:00 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.91 | 38.03 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.10 | 26.14 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 455.43 | 16.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 347.56 | 22.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 548.34 | 14.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)