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
**Last Updated:** Tue, 01 Sep 2026 07:59:28 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 256.90 | 30.14 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.58 | 27.37 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 511.69 | 15.14 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 356.04 | 21.74 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 592.95 | 13.10 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)