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
**Last Updated:** Thu, 04 Sep 2025 03:09:52 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.59 | 38.30 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 288.80 | 26.79 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 459.22 | 16.85 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 350.31 | 22.11 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.84 | 14.38 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)