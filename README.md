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
**Last Updated:** Tue, 22 Jul 2025 03:50:33 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 210.67 | 36.75 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 303.29 | 25.49 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 467.09 | 16.55 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 363.82 | 21.28 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 559.85 | 13.87 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)