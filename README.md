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
**Last Updated:** Mon, 11 Aug 2025 03:56:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 190.25 | 40.66 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 271.28 | 28.46 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 416.80 | 18.51 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 316.33 | 24.45 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 488.58 | 15.89 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)