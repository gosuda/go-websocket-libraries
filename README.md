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
**Last Updated:** Sat, 18 Jul 2026 05:11:32 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 255.60 | 30.21 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 276.59 | 27.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 504.64 | 15.34 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 341.22 | 22.65 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.0 | 573.06 | 13.56 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)