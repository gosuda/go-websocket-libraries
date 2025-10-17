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
**Last Updated:** Fri, 17 Oct 2025 03:25:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.72 | 35.24 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.90 | 26.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 445.21 | 17.36 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 346.26 | 22.36 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 537.68 | 14.45 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)