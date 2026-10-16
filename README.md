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
**Last Updated:** Fri, 16 Oct 2026 07:30:54 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 244.55 | 31.55 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 266.58 | 29.00 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 495.40 | 15.61 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 333.95 | 23.23 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 581.37 | 13.36 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)