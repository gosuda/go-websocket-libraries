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
**Last Updated:** Wed, 01 Jul 2026 06:54:33 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 265.40 | 29.17 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.51 | 26.36 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 530.28 | 14.63 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.10 | 362.08 | 21.39 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.9.1 | 602.89 | 12.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)