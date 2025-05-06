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
**Last Updated:** Tue, 06 May 2025 03:30:13 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.04 | 38.42 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.82 | 26.58 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 454.70 | 16.99 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 352.49 | 21.93 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 540.64 | 14.36 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)