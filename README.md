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
**Last Updated:** Wed, 04 Nov 2026 09:49:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 236.27 | 32.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 261.38 | 29.62 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 497.30 | 15.58 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 316.25 | 24.42 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 563.97 | 13.77 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)