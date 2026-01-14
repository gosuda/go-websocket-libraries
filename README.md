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
**Last Updated:** Wed, 14 Jan 2026 04:02:33 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.87 | 31.91 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 306.29 | 25.30 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 541.71 | 14.27 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 381.00 | 20.32 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 628.95 | 12.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)