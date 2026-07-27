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
**Last Updated:** Mon, 27 Jul 2026 06:08:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 249.92 | 30.91 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 281.12 | 27.54 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 515.30 | 15.05 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 338.07 | 22.91 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 554.20 | 14.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)