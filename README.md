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
**Last Updated:** Fri, 24 Jul 2026 05:37:17 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 261.18 | 29.58 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.26 | 27.33 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 506.59 | 15.28 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 346.00 | 22.43 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 590.82 | 13.16 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)