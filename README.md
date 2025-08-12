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
**Last Updated:** Tue, 12 Aug 2025 03:39:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.10 | 37.74 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.98 | 26.14 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.88 | 17.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 353.69 | 21.89 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 544.53 | 14.27 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)