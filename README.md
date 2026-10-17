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
**Last Updated:** Sat, 17 Oct 2026 07:11:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 262.81 | 29.48 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.18 | 26.78 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 514.79 | 15.06 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 361.86 | 21.41 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 605.58 | 12.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)