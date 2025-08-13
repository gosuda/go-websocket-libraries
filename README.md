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
**Last Updated:** Wed, 13 Aug 2025 03:41:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 197.36 | 39.16 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.88 | 26.44 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 443.87 | 17.44 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 341.14 | 22.68 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 521.21 | 14.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)