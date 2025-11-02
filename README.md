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
**Last Updated:** Sun, 02 Nov 2025 03:36:57 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 228.19 | 33.93 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.03 | 26.08 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 466.27 | 16.61 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 367.47 | 21.06 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 553.17 | 14.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)