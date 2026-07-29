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
**Last Updated:** Wed, 29 Jul 2026 05:38:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 253.00 | 30.59 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.08 | 27.66 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 499.01 | 15.49 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 344.68 | 22.46 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 584.16 | 13.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)