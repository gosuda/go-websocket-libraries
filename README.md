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
**Last Updated:** Sat, 19 Sep 2026 07:32:16 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 257.34 | 29.99 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 283.81 | 27.18 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 516.43 | 14.96 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 351.78 | 22.01 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 601.07 | 12.93 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)