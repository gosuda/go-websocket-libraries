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
**Last Updated:** Tue, 01 Jul 2025 03:50:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 196.34 | 39.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.55 | 26.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 441.97 | 17.49 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 348.44 | 22.19 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 539.22 | 14.39 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)