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
**Last Updated:** Thu, 14 May 2026 06:03:58 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.55 | 31.91 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.93 | 26.26 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 523.32 | 14.81 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.9 | 352.76 | 21.92 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 606.23 | 12.82 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)