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
**Last Updated:** Wed, 28 Jan 2026 04:00:06 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.56 | 31.80 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 301.29 | 25.66 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 553.28 | 13.98 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 384.65 | 20.11 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 646.54 | 12.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)