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
**Last Updated:** Tue, 27 Oct 2026 08:21:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 248.85 | 31.12 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 266.82 | 28.99 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 493.04 | 15.69 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 333.46 | 23.24 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 571.16 | 13.60 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)