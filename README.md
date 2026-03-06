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
**Last Updated:** Fri, 06 Mar 2026 04:22:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 234.14 | 33.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 281.91 | 27.40 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 513.07 | 15.09 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 352.93 | 21.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 584.31 | 13.30 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)