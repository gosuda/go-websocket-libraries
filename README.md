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
**Last Updated:** Wed, 13 May 2026 06:04:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.00 | 32.01 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.96 | 26.38 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 523.61 | 14.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 355.11 | 21.80 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 597.97 | 12.97 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)