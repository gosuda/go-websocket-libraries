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
**Last Updated:** Fri, 02 Oct 2026 08:59:24 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 242.95 | 31.87 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.68 | 27.30 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 477.61 | 16.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.0 | 343.92 | 22.46 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 593.68 | 13.08 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)