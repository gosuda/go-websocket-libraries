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
**Last Updated:** Tue, 14 Jul 2026 05:16:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 275.55 | 28.20 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.65 | 25.90 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 523.49 | 14.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.11 | 369.38 | 20.99 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 601.34 | 12.94 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)