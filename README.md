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
**Last Updated:** Sat, 04 Apr 2026 04:31:46 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.18 | 31.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.89 | 26.80 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 517.13 | 14.96 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 350.84 | 22.05 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 605.80 | 12.82 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)