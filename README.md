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
**Last Updated:** Thu, 05 Feb 2026 04:28:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 249.14 | 31.04 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 319.10 | 24.28 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 561.26 | 13.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 394.56 | 19.64 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 643.30 | 12.09 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)