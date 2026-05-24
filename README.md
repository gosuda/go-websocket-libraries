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
**Last Updated:** Sun, 24 May 2026 06:24:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 253.31 | 30.57 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 312.88 | 24.76 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 550.03 | 14.08 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 384.49 | 20.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 635.20 | 12.24 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)