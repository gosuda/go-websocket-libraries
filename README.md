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
**Last Updated:** Sat, 25 Jul 2026 05:31:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 258.45 | 30.02 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 280.95 | 27.60 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 501.50 | 15.44 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 354.51 | 21.86 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 587.03 | 13.27 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)