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
**Last Updated:** Sat, 25 Apr 2026 05:01:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 241.22 | 32.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.61 | 26.91 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 533.43 | 14.52 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 350.64 | 22.08 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 608.71 | 12.77 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)