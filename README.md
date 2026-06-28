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
**Last Updated:** Sun, 28 Jun 2026 06:46:34 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 266.24 | 29.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.32 | 26.05 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 504.01 | 15.36 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 359.09 | 21.56 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 597.57 | 13.02 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)