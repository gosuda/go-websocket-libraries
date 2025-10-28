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
**Last Updated:** Tue, 28 Oct 2025 03:32:12 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 199.81 | 38.68 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 259.65 | 29.80 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 382.76 | 20.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 315.16 | 24.55 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 479.08 | 16.19 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)