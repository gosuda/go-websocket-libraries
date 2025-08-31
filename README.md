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
**Last Updated:** Sun, 31 Aug 2025 03:28:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 201.39 | 38.36 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.55 | 25.90 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.97 | 17.04 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 351.85 | 21.98 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 535.01 | 14.52 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)