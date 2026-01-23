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
**Last Updated:** Fri, 23 Jan 2026 03:56:58 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.20 | 31.73 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 295.39 | 26.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 549.32 | 14.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 372.97 | 20.76 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 640.24 | 12.16 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)