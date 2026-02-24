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
**Last Updated:** Tue, 24 Feb 2026 04:32:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.05 | 32.00 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 304.71 | 25.44 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 541.55 | 14.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 374.67 | 20.67 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 636.51 | 12.22 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)