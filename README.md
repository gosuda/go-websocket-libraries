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
**Last Updated:** Tue, 04 Nov 2025 03:33:15 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 206.75 | 37.41 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.26 | 26.56 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 419.38 | 18.45 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 329.90 | 23.47 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 527.20 | 14.73 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)