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
**Last Updated:** Sat, 29 Nov 2025 03:35:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 217.57 | 35.56 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.82 | 26.85 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 451.01 | 17.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 348.83 | 22.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 544.93 | 14.24 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)