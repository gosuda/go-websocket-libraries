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
**Last Updated:** Tue, 29 Jul 2025 04:01:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 191.01 | 40.53 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.06 | 28.03 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 422.89 | 18.29 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 332.98 | 23.23 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 512.88 | 15.13 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)