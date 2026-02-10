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
**Last Updated:** Tue, 10 Feb 2026 04:55:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 243.54 | 31.84 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 292.68 | 26.43 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 548.47 | 14.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 379.01 | 20.43 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 632.62 | 12.30 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)