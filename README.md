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
**Last Updated:** Wed, 22 Jul 2026 05:36:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 242.38 | 31.91 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.98 | 27.92 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 490.17 | 15.79 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 333.13 | 23.22 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 567.60 | 13.71 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)