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
**Last Updated:** Mon, 13 Jul 2026 05:54:37 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 340.34 | 22.78 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 372.36 | 20.83 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 641.68 | 12.10 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.11 | 455.87 | 17.01 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 757.59 | 10.27 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)