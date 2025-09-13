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
**Last Updated:** Sat, 13 Sep 2025 03:03:32 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 221.60 | 34.98 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 289.16 | 26.74 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 446.56 | 17.34 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 350.82 | 22.10 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 545.18 | 14.26 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)