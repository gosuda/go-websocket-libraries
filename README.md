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
**Last Updated:** Sat, 15 Nov 2025 03:30:35 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.18 | 35.39 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.45 | 26.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.53 | 17.15 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 346.70 | 22.32 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 534.03 | 14.56 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)