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
**Last Updated:** Wed, 09 Sep 2026 07:34:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 266.47 | 29.10 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 297.24 | 26.04 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 516.22 | 14.99 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 362.40 | 21.36 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.1 | 608.21 | 12.77 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)