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
**Last Updated:** Tue, 13 Oct 2026 09:44:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 324.70 | 23.88 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 340.43 | 22.82 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 608.66 | 12.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 421.15 | 18.43 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 726.88 | 10.70 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)