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
**Last Updated:** Sat, 30 Aug 2025 03:18:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 197.13 | 39.20 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.45 | 26.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.81 | 17.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 354.54 | 21.85 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 549.39 | 14.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)