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
**Last Updated:** Tue, 16 Jun 2026 08:33:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 264.27 | 29.32 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.70 | 26.40 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 528.96 | 14.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 360.05 | 21.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 607.88 | 12.78 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)