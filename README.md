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
**Last Updated:** Sun, 19 Apr 2026 05:18:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 236.80 | 32.73 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.10 | 26.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 516.20 | 15.01 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 353.78 | 21.89 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 614.05 | 12.64 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)