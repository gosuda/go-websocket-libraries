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
**Last Updated:** Sun, 06 Sep 2026 07:18:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 262.72 | 29.44 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.23 | 26.15 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 504.27 | 15.35 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 361.22 | 21.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 600.19 | 12.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)