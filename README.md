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
**Last Updated:** Tue, 18 Aug 2026 03:26:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 316.17 | 24.50 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 382.56 | 20.28 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 645.20 | 12.03 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 448.54 | 17.32 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 803.02 | 9.69 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)