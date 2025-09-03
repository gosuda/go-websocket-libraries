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
**Last Updated:** Wed, 03 Sep 2025 03:09:38 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 192.84 | 39.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 271.20 | 28.49 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 435.73 | 17.73 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 329.99 | 23.44 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 516.41 | 15.05 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)