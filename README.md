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
**Last Updated:** Sun, 31 May 2026 06:43:42 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 239.19 | 32.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.06 | 26.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 538.58 | 14.39 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 363.27 | 21.31 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 618.97 | 12.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)