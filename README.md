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
**Last Updated:** Thu, 24 Sep 2026 07:48:40 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 258.64 | 29.86 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 296.44 | 26.10 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 506.54 | 15.29 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.0 | 353.52 | 21.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 592.82 | 13.09 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)