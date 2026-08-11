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
**Last Updated:** Tue, 11 Aug 2026 04:05:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 261.05 | 29.65 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.12 | 26.89 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 494.84 | 15.64 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 353.25 | 21.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 585.12 | 13.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)