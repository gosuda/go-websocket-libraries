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
**Last Updated:** Tue, 30 Dec 2025 03:52:08 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 263.04 | 29.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 358.64 | 21.63 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 606.27 | 12.77 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 427.85 | 18.13 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 774.32 | 10.06 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)