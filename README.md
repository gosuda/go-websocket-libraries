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
**Last Updated:** Sun, 01 Nov 2026 09:41:32 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 236.95 | 32.63 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 270.97 | 28.54 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 493.95 | 15.65 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 329.44 | 23.53 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 565.28 | 13.75 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)