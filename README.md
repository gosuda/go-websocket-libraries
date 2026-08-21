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
**Last Updated:** Fri, 21 Aug 2026 03:33:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 262.71 | 29.50 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.46 | 26.41 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 508.75 | 15.22 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 360.99 | 21.45 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 589.19 | 13.19 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)