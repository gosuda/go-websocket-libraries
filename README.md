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
**Last Updated:** Sun, 25 Jan 2026 04:07:13 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 238.93 | 32.43 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.68 | 26.40 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 547.38 | 14.18 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 381.04 | 20.32 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 636.80 | 12.20 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)