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
**Last Updated:** Sun, 05 Jul 2026 06:26:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 266.60 | 29.08 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 300.16 | 25.76 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 526.08 | 14.72 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.11 | 362.52 | 21.38 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 598.78 | 12.98 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)