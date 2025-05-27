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
**Last Updated:** Tue, 27 May 2025 03:33:27 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 199.98 | 38.60 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.21 | 26.58 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 448.96 | 17.19 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 345.84 | 22.34 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.54 | 14.36 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)