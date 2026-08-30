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
**Last Updated:** Sun, 30 Aug 2026 08:32:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 249.35 | 31.00 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 272.36 | 28.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 496.84 | 15.58 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 343.97 | 22.49 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 579.86 | 13.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)