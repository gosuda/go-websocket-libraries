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
**Last Updated:** Mon, 13 Oct 2025 03:33:19 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 253.35 | 30.58 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 373.53 | 20.75 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 516.35 | 14.99 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 432.16 | 17.94 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 680.36 | 11.43 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)