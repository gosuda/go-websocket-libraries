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
**Last Updated:** Fri, 27 Mar 2026 04:58:10 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 231.15 | 33.46 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.54 | 27.67 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 519.30 | 14.93 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 349.31 | 22.19 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 591.12 | 13.10 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)