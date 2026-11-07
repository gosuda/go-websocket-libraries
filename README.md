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
**Last Updated:** Sat, 07 Nov 2026 09:12:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 443.90 | 17.49 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 481.66 | 16.15 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 848.54 | 9.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 584.01 | 13.31 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 1036.79 | 7.51 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)