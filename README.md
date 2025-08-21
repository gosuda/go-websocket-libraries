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
**Last Updated:** Thu, 21 Aug 2025 03:31:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 197.19 | 39.25 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.52 | 27.30 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 448.28 | 17.25 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 345.45 | 22.40 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 543.35 | 14.32 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)