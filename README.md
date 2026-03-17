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
**Last Updated:** Tue, 17 Mar 2026 04:31:16 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 229.41 | 33.64 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 275.57 | 28.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 489.55 | 15.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 344.33 | 22.51 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 576.68 | 13.48 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)