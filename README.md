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
**Last Updated:** Sun, 25 Oct 2026 07:48:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 449.08 | 17.30 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 489.12 | 15.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 857.80 | 9.05 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 586.14 | 13.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 998.99 | 7.79 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)