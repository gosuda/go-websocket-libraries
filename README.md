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
**Last Updated:** Sat, 17 Jan 2026 03:46:42 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 241.70 | 32.04 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.52 | 25.91 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 548.82 | 14.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 387.97 | 19.96 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 655.26 | 11.88 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)