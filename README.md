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
**Last Updated:** Mon, 07 Jul 2025 03:45:42 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 196.20 | 39.37 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.89 | 27.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 435.22 | 17.75 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 335.21 | 23.09 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 513.89 | 15.11 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)