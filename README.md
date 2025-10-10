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
**Last Updated:** Fri, 10 Oct 2025 03:23:11 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.74 | 34.26 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.83 | 26.25 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 458.37 | 16.86 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 359.22 | 21.56 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 559.68 | 13.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)