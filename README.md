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
**Last Updated:** Sun, 03 Aug 2025 04:03:05 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 208.35 | 37.14 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.98 | 26.04 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 467.10 | 16.54 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 363.03 | 21.32 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.45 | 14.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)