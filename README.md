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
**Last Updated:** Thu, 02 Oct 2025 03:19:25 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 217.06 | 35.64 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 287.84 | 26.88 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 450.44 | 17.17 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 343.05 | 22.57 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 538.39 | 14.43 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)