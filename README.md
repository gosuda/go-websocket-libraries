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
**Last Updated:** Mon, 14 Sep 2026 08:06:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 259.45 | 29.82 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.33 | 27.14 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 509.79 | 15.19 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 355.23 | 21.83 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 593.33 | 13.09 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)