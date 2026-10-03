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
**Last Updated:** Sat, 03 Oct 2026 08:33:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 330.48 | 23.46 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 402.94 | 19.28 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 668.16 | 11.61 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.0 | 475.31 | 16.35 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 851.49 | 9.14 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)