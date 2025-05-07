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
**Last Updated:** Wed, 07 May 2025 03:32:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 197.85 | 39.05 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.95 | 26.86 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 430.67 | 17.91 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 341.48 | 22.54 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.8 | 523.12 | 14.84 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)