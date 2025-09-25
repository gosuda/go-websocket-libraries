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
**Last Updated:** Thu, 25 Sep 2025 03:24:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 221.95 | 34.85 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.54 | 26.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 452.82 | 17.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 355.38 | 21.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 543.61 | 14.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)