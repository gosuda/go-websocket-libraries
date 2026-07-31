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
**Last Updated:** Fri, 31 Jul 2026 05:52:12 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 497.96 | 15.61 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 645.25 | 12.03 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 1044.92 | 7.43 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 717.56 | 10.84 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 1321.08 | 5.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)