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
**Last Updated:** Wed, 06 May 2026 05:45:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 239.09 | 32.37 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.47 | 26.23 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 512.61 | 15.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 346.54 | 22.37 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 594.65 | 13.07 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)