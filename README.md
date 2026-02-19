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
**Last Updated:** Thu, 19 Feb 2026 04:33:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 238.80 | 32.45 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.88 | 26.29 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 522.14 | 14.84 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 365.38 | 21.23 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 599.44 | 12.97 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)