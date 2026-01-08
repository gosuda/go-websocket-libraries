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
**Last Updated:** Thu, 08 Jan 2026 03:53:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 231.54 | 33.44 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.50 | 26.67 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 523.44 | 14.80 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 359.74 | 21.52 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 624.02 | 12.46 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)