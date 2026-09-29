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
**Last Updated:** Tue, 29 Sep 2026 09:00:51 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 238.54 | 32.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 261.64 | 29.53 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 461.86 | 16.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.0 | 314.88 | 24.55 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 531.46 | 14.62 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)