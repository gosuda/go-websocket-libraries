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
**Last Updated:** Sat, 14 Nov 2026 09:23:08 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 500.80 | 15.51 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 638.68 | 12.17 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 1027.33 | 7.56 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 714.52 | 10.88 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 1319.49 | 5.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)