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
**Last Updated:** Sun, 08 Nov 2026 09:27:59 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 253.65 | 30.46 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.5.0 | 282.66 | 27.39 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 514.25 | 15.04 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.2 | 340.45 | 22.73 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 605.37 | 12.81 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)