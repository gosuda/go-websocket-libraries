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
**Last Updated:** Thu, 25 Jun 2026 06:34:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 347.25 | 22.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 381.99 | 20.31 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 646.50 | 11.99 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 471.10 | 16.49 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 762.18 | 10.21 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)