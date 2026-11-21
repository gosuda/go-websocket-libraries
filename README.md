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
**Last Updated:** Sat, 21 Nov 2026 10:09:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 256.16 | 30.22 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 289.48 | 26.75 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 478.10 | 16.19 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 354.72 | 21.81 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 591.46 | 13.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)