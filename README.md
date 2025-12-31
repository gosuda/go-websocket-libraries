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
**Last Updated:** Wed, 31 Dec 2025 03:51:30 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 254.63 | 30.40 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 308.56 | 25.12 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 560.40 | 13.82 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 394.91 | 19.61 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 626.20 | 12.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)