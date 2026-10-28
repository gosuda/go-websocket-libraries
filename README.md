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
**Last Updated:** Wed, 28 Oct 2026 08:58:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 400.26 | 19.43 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 405.55 | 19.17 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 786.92 | 9.86 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 531.23 | 14.63 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 905.55 | 8.61 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)