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
**Last Updated:** Mon, 23 Nov 2026 11:13:14 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 317.26 | 24.42 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 379.02 | 20.47 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 641.76 | 12.08 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 421.72 | 18.40 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 791.30 | 9.83 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)