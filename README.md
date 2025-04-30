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
**Last Updated:** Wed, 30 Apr 2025 02:07:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 206.85 | 37.35 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.84 | 27.91 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.25 | 17.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 356.67 | 21.69 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 545.79 | 14.21 |
<!-- BENCHMARK_TABLE_END -->



**Performance Over Time (Connections/Second):**

![Benchmark Performance Graph](benchmark_performance.png)