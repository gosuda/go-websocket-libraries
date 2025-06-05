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
**Last Updated:** Thu, 05 Jun 2025 03:38:06 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.10 | 38.24 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.47 | 26.15 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 450.43 | 17.16 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 352.56 | 21.92 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.56 | 14.38 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)