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
**Last Updated:** Tue, 14 Oct 2025 03:23:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 229.32 | 33.78 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.66 | 25.84 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 471.49 | 16.40 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 366.17 | 21.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 561.78 | 13.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)