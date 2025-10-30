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
**Last Updated:** Thu, 30 Oct 2025 03:33:58 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 224.36 | 34.50 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.71 | 27.08 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 461.30 | 16.74 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 354.88 | 21.82 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 551.51 | 14.10 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)