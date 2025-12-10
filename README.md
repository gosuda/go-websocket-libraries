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
**Last Updated:** Wed, 10 Dec 2025 03:46:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 225.01 | 34.41 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.57 | 26.43 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 458.54 | 16.87 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 357.27 | 21.68 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 546.45 | 14.23 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)