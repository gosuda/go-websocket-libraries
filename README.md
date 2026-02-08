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
**Last Updated:** Sun, 08 Feb 2026 04:54:21 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 231.85 | 33.34 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.06 | 26.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 534.93 | 14.44 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 362.91 | 21.33 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 616.00 | 12.62 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)