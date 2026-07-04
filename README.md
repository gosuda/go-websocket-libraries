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
**Last Updated:** Sat, 04 Jul 2026 06:01:13 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 289.28 | 26.83 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 341.54 | 22.68 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 555.93 | 13.98 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.11 | 393.33 | 19.74 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 692.41 | 11.21 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)