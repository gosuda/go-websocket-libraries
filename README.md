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
**Last Updated:** Fri, 06 Feb 2026 04:29:02 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 244.23 | 31.69 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.70 | 25.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 556.74 | 13.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 382.44 | 20.24 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 637.17 | 12.20 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)