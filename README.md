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
**Last Updated:** Mon, 04 May 2026 05:51:11 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.44 | 31.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.34 | 26.10 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 522.70 | 14.82 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 356.44 | 21.72 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 615.71 | 12.64 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)