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
**Last Updated:** Thu, 20 Aug 2026 04:30:32 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 249.23 | 31.02 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 272.69 | 28.45 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 516.02 | 15.03 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 334.76 | 23.08 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 593.20 | 13.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)