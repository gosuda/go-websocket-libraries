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
**Last Updated:** Mon, 09 Nov 2026 10:17:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 246.52 | 31.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 274.71 | 28.07 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 485.15 | 15.93 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 331.55 | 23.35 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 576.48 | 13.48 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)