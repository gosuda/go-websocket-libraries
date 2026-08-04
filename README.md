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
**Last Updated:** Tue, 04 Aug 2026 05:33:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 246.67 | 31.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 275.71 | 28.03 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 475.92 | 16.22 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 334.41 | 23.16 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 548.86 | 14.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)