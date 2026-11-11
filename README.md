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
**Last Updated:** Wed, 11 Nov 2026 09:49:48 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 338.62 | 22.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 353.07 | 21.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 618.91 | 12.54 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 437.19 | 17.74 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 748.22 | 10.39 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)