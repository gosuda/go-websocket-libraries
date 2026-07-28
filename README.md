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
**Last Updated:** Tue, 28 Jul 2026 05:34:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 324.23 | 23.93 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 367.04 | 21.14 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 628.10 | 12.37 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 429.44 | 18.07 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 740.77 | 10.51 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)