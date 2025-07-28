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
**Last Updated:** Mon, 28 Jul 2025 03:59:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 203.55 | 37.98 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.44 | 26.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 459.20 | 16.84 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 354.61 | 21.84 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 549.85 | 14.12 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)