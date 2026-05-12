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
**Last Updated:** Tue, 12 May 2026 05:55:10 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 275.22 | 28.15 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 379.88 | 20.41 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 613.32 | 12.62 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 423.28 | 18.30 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 757.55 | 10.28 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)