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
**Last Updated:** Mon, 01 Jun 2026 07:31:53 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 247.06 | 31.36 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.75 | 26.04 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 527.16 | 14.65 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 366.35 | 21.13 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 634.23 | 12.25 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)