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
**Last Updated:** Sun, 15 Feb 2026 04:36:08 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.32 | 31.76 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.33 | 26.11 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 523.26 | 14.79 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 371.29 | 20.83 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 604.85 | 12.83 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)