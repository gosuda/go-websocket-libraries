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
**Last Updated:** Tue, 09 Sep 2025 03:24:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 231.53 | 33.44 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 305.96 | 25.31 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 477.72 | 16.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 369.22 | 20.96 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 573.25 | 13.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)