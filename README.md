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
**Last Updated:** Wed, 25 Jun 2025 03:41:01 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 202.70 | 38.11 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.53 | 26.48 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 435.15 | 17.74 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 358.43 | 21.58 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 529.96 | 14.62 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)