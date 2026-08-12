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
**Last Updated:** Wed, 12 Aug 2026 04:26:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 267.33 | 28.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 303.01 | 25.57 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 522.28 | 14.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 362.72 | 21.33 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 609.77 | 12.76 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)