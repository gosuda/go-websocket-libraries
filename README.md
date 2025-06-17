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
**Last Updated:** Tue, 17 Jun 2025 03:38:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.49 | 38.10 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.83 | 26.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.09 | 17.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 350.85 | 22.02 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 545.67 | 14.24 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)