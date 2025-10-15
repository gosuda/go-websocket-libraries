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
**Last Updated:** Wed, 15 Oct 2025 03:27:56 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 213.83 | 36.19 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 278.26 | 27.82 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 428.20 | 18.06 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 344.44 | 22.49 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 538.34 | 14.42 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)