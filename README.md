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
**Last Updated:** Tue, 10 Mar 2026 04:24:32 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 223.92 | 34.60 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 261.91 | 29.64 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 479.47 | 16.14 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 338.51 | 22.88 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 562.36 | 13.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)