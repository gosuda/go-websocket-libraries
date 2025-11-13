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
**Last Updated:** Thu, 13 Nov 2025 03:37:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 219.84 | 35.18 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.06 | 26.85 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 442.15 | 17.49 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 356.91 | 21.68 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 540.01 | 14.41 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)