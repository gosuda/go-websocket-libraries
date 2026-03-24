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
**Last Updated:** Tue, 24 Mar 2026 04:33:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 235.25 | 32.88 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.44 | 27.40 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 518.34 | 14.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 349.73 | 22.17 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.0 | 583.15 | 13.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)