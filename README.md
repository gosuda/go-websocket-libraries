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
**Last Updated:** Wed, 17 Jun 2026 07:38:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 264.04 | 29.30 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 299.14 | 25.89 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 488.70 | 15.84 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 363.86 | 21.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 593.82 | 13.08 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)