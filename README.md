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
**Last Updated:** Sat, 11 Jul 2026 05:28:58 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 262.35 | 29.44 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.31 | 26.74 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 508.90 | 15.22 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.11 | 357.28 | 21.66 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 591.63 | 13.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)