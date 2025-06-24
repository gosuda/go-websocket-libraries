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
**Last Updated:** Tue, 24 Jun 2025 03:39:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 192.82 | 40.09 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.37 | 27.69 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 426.03 | 18.06 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 335.69 | 23.04 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 519.41 | 14.97 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)