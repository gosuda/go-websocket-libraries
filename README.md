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
**Last Updated:** Mon, 20 Oct 2025 03:38:34 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.92 | 35.18 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 299.88 | 25.80 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 454.95 | 17.02 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 355.88 | 21.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 543.37 | 14.29 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)