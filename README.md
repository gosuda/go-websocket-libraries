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
**Last Updated:** Sat, 14 Feb 2026 04:24:24 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 245.13 | 31.57 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.09 | 26.51 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 526.70 | 14.70 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 372.46 | 20.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 600.52 | 12.94 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)