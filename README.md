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
**Last Updated:** Thu, 01 May 2025 03:37:12 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 198.46 | 38.82 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.36 | 27.77 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 434.03 | 17.77 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 336.70 | 22.97 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 515.28 | 15.10 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)