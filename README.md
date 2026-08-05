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
**Last Updated:** Wed, 05 Aug 2026 05:33:13 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 276.08 | 28.00 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 305.60 | 25.34 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 528.58 | 14.63 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 373.39 | 20.75 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 615.39 | 12.62 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)