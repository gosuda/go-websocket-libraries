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
**Last Updated:** Tue, 02 Dec 2025 03:43:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.92 | 34.27 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.86 | 26.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 459.58 | 16.82 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 358.59 | 21.60 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 568.05 | 13.70 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)