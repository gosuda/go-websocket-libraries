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
**Last Updated:** Mon, 03 Nov 2025 03:40:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 221.22 | 34.98 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 284.80 | 27.11 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 450.90 | 17.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 349.08 | 22.16 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 537.66 | 14.46 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)