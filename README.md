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
**Last Updated:** Wed, 23 Sep 2026 07:58:12 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 337.24 | 23.00 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 361.84 | 21.45 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 636.60 | 12.18 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 451.86 | 17.18 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 753.51 | 10.33 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)