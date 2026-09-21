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
**Last Updated:** Mon, 21 Sep 2026 08:12:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 247.48 | 31.26 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 274.05 | 28.21 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 499.46 | 15.49 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 337.98 | 22.88 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 584.92 | 13.26 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)