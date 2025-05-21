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
**Last Updated:** Wed, 21 May 2025 03:32:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.09 | 37.50 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 301.71 | 25.64 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 458.32 | 16.88 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 368.35 | 20.96 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 554.43 | 14.02 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)