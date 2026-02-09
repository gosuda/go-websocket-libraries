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
**Last Updated:** Mon, 09 Feb 2026 04:40:07 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 301.84 | 25.71 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 420.23 | 18.46 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 676.17 | 11.50 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 498.40 | 15.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 861.97 | 9.03 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)