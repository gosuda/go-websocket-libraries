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
**Last Updated:** Tue, 16 Sep 2025 03:18:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 218.79 | 35.39 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.40 | 26.41 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 449.34 | 17.20 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 349.03 | 22.20 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 531.60 | 14.62 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)