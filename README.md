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
**Last Updated:** Mon, 02 Nov 2026 09:56:29 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 315.05 | 24.66 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 376.38 | 20.60 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 641.70 | 12.08 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 435.63 | 17.82 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 816.34 | 9.54 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)