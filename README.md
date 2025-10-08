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
**Last Updated:** Wed, 08 Oct 2025 03:19:45 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 239.61 | 32.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 352.25 | 22.04 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 480.81 | 16.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 396.07 | 19.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 623.29 | 12.49 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)