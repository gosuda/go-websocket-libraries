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
**Last Updated:** Thu, 18 Sep 2025 03:19:52 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.75 | 35.08 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 291.57 | 26.52 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 446.71 | 17.30 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 349.01 | 22.16 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 521.28 | 14.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)