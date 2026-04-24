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
**Last Updated:** Fri, 24 Apr 2026 05:23:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 251.09 | 30.93 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.39 | 26.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 543.45 | 14.25 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 363.24 | 21.30 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 630.93 | 12.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)