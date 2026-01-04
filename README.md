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
**Last Updated:** Sun, 04 Jan 2026 04:04:43 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 237.45 | 32.59 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 290.24 | 26.69 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 542.44 | 14.26 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 362.86 | 21.34 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 622.70 | 12.49 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)