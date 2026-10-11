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
**Last Updated:** Sun, 11 Oct 2026 09:17:17 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 362.42 | 21.42 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 398.11 | 19.50 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 741.03 | 10.48 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 499.96 | 15.53 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 856.34 | 9.09 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)