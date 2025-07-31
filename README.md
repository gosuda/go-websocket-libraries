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
**Last Updated:** Thu, 31 Jul 2025 03:54:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.59 | 38.27 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.19 | 26.32 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 456.77 | 16.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 346.19 | 22.36 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 535.10 | 14.49 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)