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
**Last Updated:** Fri, 16 May 2025 03:34:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.43 | 38.39 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.79 | 26.91 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 452.79 | 17.07 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 350.22 | 22.08 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.24 | 14.40 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)