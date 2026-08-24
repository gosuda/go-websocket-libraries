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
**Last Updated:** Mon, 24 Aug 2026 03:35:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 452.29 | 17.18 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 491.17 | 15.83 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 841.90 | 9.22 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 592.24 | 13.13 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 1009.90 | 7.71 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)