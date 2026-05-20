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
**Last Updated:** Wed, 20 May 2026 06:35:13 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.69 | 31.91 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.01 | 26.00 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 539.33 | 14.37 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 368.18 | 21.04 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 621.74 | 12.51 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)