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
**Last Updated:** Fri, 09 Jan 2026 03:53:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.39 | 31.93 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.93 | 26.08 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 542.85 | 14.25 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 373.58 | 20.72 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 637.56 | 12.19 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)