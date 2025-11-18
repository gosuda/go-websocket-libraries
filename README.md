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
**Last Updated:** Tue, 18 Nov 2025 03:36:08 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.78 | 34.24 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.11 | 26.19 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 469.78 | 16.49 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 359.38 | 21.54 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 565.46 | 13.75 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)