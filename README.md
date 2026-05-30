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
**Last Updated:** Sat, 30 May 2026 06:03:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 235.43 | 32.88 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.55 | 27.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 522.77 | 14.84 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 347.50 | 22.31 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 600.47 | 12.96 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)