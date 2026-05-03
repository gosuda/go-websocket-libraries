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
**Last Updated:** Sun, 03 May 2026 05:45:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.33 | 31.79 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.51 | 25.75 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 514.18 | 15.05 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 358.29 | 21.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 589.73 | 13.20 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)