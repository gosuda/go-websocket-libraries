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
**Last Updated:** Tue, 31 Mar 2026 05:00:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 245.33 | 31.50 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 298.43 | 25.92 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 527.64 | 14.66 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 362.63 | 21.40 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 624.17 | 12.45 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)