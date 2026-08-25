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
**Last Updated:** Tue, 25 Aug 2026 03:29:22 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 470.59 | 16.52 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 581.65 | 13.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 934.31 | 8.31 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 661.60 | 11.75 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 1163.27 | 6.69 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)