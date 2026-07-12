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
**Last Updated:** Sun, 12 Jul 2026 05:44:27 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 260.02 | 29.78 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.31 | 26.94 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 519.11 | 14.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.11 | 351.75 | 22.01 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 604.88 | 12.85 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)