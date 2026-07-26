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
**Last Updated:** Sun, 26 Jul 2026 05:47:46 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 249.78 | 30.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.78 | 27.97 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 482.28 | 16.05 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 333.66 | 23.20 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 555.78 | 13.97 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)