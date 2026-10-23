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
**Last Updated:** Fri, 23 Oct 2026 07:55:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 363.35 | 21.37 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 436.62 | 17.80 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 701.67 | 11.08 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 532.57 | 14.60 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 888.56 | 8.77 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)