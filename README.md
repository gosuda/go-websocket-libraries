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
**Last Updated:** Sat, 07 Feb 2026 04:21:34 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.15 | 35.34 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 302.93 | 25.53 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 514.40 | 15.04 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 325.12 | 23.79 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 735.83 | 10.59 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)