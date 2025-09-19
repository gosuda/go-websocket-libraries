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
**Last Updated:** Fri, 19 Sep 2025 03:23:23 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 215.33 | 35.92 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.44 | 27.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 445.87 | 17.34 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 347.70 | 22.23 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 520.90 | 14.92 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)