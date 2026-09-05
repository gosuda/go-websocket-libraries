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
**Last Updated:** Sat, 05 Sep 2026 07:07:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 255.17 | 30.27 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 285.87 | 27.08 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 513.01 | 15.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 357.22 | 21.70 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 607.49 | 12.80 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)