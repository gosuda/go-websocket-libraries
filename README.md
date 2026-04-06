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
**Last Updated:** Mon, 06 Apr 2026 05:13:09 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 242.39 | 31.94 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 294.25 | 26.29 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 535.21 | 14.47 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.8 | 353.07 | 21.94 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 614.09 | 12.65 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)