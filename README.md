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
**Last Updated:** Wed, 14 Oct 2026 08:00:27 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 341.35 | 22.70 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 418.82 | 18.52 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 694.74 | 11.16 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 485.39 | 15.98 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 911.17 | 8.54 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)