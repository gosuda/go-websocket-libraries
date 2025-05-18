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
**Last Updated:** Sun, 18 May 2025 03:36:34 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 208.68 | 37.05 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.74 | 25.86 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 471.55 | 16.39 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 362.64 | 21.32 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 564.10 | 13.76 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)