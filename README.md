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
**Last Updated:** Tue, 02 Jun 2026 07:10:55 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 244.19 | 31.75 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.87 | 26.49 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 510.08 | 15.18 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 358.91 | 21.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 616.26 | 12.62 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)