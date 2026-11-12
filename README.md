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
**Last Updated:** Thu, 12 Nov 2026 10:15:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 250.07 | 30.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 278.32 | 27.78 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 495.64 | 15.61 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 335.54 | 23.06 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 585.78 | 13.27 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)