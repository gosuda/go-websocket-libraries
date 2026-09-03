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
**Last Updated:** Thu, 03 Sep 2026 07:23:10 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 254.82 | 30.38 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.54 | 27.96 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 518.64 | 14.92 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 346.91 | 22.29 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 589.61 | 13.19 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)