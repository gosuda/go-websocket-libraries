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
**Last Updated:** Sat, 16 Aug 2025 03:37:06 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 207.63 | 37.31 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 311.61 | 24.84 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 466.90 | 16.57 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 365.85 | 21.19 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 569.35 | 13.65 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)