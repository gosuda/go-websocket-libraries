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
**Last Updated:** Fri, 10 Apr 2026 05:14:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 250.98 | 30.89 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.89 | 26.47 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 500.74 | 15.50 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 367.01 | 21.10 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 568.58 | 13.70 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)