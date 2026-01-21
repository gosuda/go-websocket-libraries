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
**Last Updated:** Wed, 21 Jan 2026 03:58:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 224.39 | 34.48 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 270.24 | 28.57 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.82 | 15.02 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 343.03 | 22.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 583.16 | 13.32 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)