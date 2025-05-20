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
**Last Updated:** Tue, 20 May 2025 03:33:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 188.89 | 40.85 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 272.39 | 28.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 423.76 | 18.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 318.74 | 24.29 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 494.13 | 15.68 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)