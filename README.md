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
**Last Updated:** Sat, 01 Aug 2026 05:42:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 253.27 | 30.42 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 272.76 | 28.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 495.55 | 15.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 335.40 | 23.06 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 572.96 | 13.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)