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
**Last Updated:** Sat, 24 May 2025 03:28:52 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.84 | 38.29 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.35 | 26.16 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.31 | 17.06 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 362.98 | 21.30 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 532.64 | 14.60 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)