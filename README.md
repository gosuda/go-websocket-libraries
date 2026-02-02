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
**Last Updated:** Mon, 02 Feb 2026 04:38:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 224.42 | 34.44 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.16 | 27.81 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 511.76 | 15.13 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 345.64 | 22.38 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 586.12 | 13.28 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)