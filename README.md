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
**Last Updated:** Mon, 20 Jul 2026 05:58:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 265.01 | 29.18 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.35 | 26.93 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.26 | 15.01 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 352.90 | 21.93 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 612.18 | 12.71 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)