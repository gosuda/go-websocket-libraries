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
**Last Updated:** Fri, 20 Nov 2026 10:26:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 508.16 | 15.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 661.03 | 11.76 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 1067.31 | 7.27 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 734.56 | 10.59 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 1347.21 | 5.78 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)