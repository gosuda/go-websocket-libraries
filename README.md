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
**Last Updated:** Mon, 26 Oct 2026 08:06:18 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 242.24 | 31.96 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 264.34 | 29.27 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 487.37 | 15.87 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 325.43 | 23.78 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 564.34 | 13.76 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)