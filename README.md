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
**Last Updated:** Wed, 21 Oct 2026 07:59:33 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 301.29 | 25.73 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 372.27 | 20.82 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 621.01 | 12.47 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 428.76 | 18.07 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 782.14 | 9.94 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)