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
**Last Updated:** Sun, 07 Dec 2025 03:48:35 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 238.66 | 32.48 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 358.46 | 21.60 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 497.60 | 15.56 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 405.35 | 19.14 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 653.95 | 11.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)