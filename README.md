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
**Last Updated:** Fri, 22 Aug 2025 03:31:44 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 192.11 | 40.34 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.46 | 26.47 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 435.22 | 17.77 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 335.15 | 23.11 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 533.97 | 14.55 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)