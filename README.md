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
**Last Updated:** Fri, 03 Apr 2026 04:55:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 244.39 | 31.69 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.18 | 26.30 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 539.07 | 14.34 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 359.54 | 21.54 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 613.87 | 12.65 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)