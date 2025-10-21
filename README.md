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
**Last Updated:** Tue, 21 Oct 2025 03:29:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 211.93 | 36.41 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.24 | 27.74 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 428.21 | 18.03 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 342.28 | 22.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 531.14 | 14.64 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)