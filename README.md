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
**Last Updated:** Fri, 22 May 2026 06:34:05 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 282.83 | 27.45 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 350.66 | 22.13 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 597.35 | 12.98 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 412.12 | 18.83 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 672.86 | 11.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)