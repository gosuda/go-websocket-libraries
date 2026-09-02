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
**Last Updated:** Wed, 02 Sep 2026 07:19:03 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 416.24 | 18.65 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 503.81 | 15.40 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 830.74 | 9.35 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 593.59 | 13.09 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 1028.27 | 7.57 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)