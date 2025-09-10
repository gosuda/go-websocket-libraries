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
**Last Updated:** Wed, 10 Sep 2025 03:10:10 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 220.59 | 35.07 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.50 | 26.60 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 459.68 | 16.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 351.06 | 22.09 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 559.28 | 13.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)