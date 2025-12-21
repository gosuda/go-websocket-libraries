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
**Last Updated:** Sun, 21 Dec 2025 03:51:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 254.24 | 30.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 344.14 | 22.52 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 585.84 | 13.21 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 408.98 | 18.96 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 726.82 | 10.70 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)