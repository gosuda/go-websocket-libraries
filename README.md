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
**Last Updated:** Wed, 04 Feb 2026 04:22:46 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 284.98 | 27.22 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 374.72 | 20.64 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 652.97 | 11.89 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 444.40 | 17.45 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 833.51 | 9.34 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)