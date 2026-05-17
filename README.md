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
**Last Updated:** Sun, 17 May 2026 06:04:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.63 | 31.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.11 | 26.64 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 519.31 | 14.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 372.07 | 20.84 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 613.33 | 12.68 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)