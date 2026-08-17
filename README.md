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
**Last Updated:** Mon, 17 Aug 2026 03:30:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 257.09 | 30.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.85 | 27.68 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 521.54 | 14.87 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 353.39 | 21.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 594.96 | 13.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)