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
**Last Updated:** Mon, 07 Sep 2026 07:32:20 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 249.70 | 31.00 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 279.54 | 27.65 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 503.43 | 15.36 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 342.11 | 22.64 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 579.54 | 13.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)