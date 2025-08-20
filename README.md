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
**Last Updated:** Wed, 20 Aug 2025 03:31:24 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 196.34 | 39.43 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.20 | 26.77 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 456.71 | 16.94 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 343.81 | 22.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 541.85 | 14.36 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)