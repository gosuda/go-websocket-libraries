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
**Last Updated:** Tue, 03 Mar 2026 04:28:34 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.76 | 31.80 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 335.28 | 23.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 533.14 | 14.57 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 369.14 | 20.98 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 639.78 | 12.16 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)