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
**Last Updated:** Wed, 30 Jul 2025 03:55:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 193.68 | 39.75 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.32 | 27.87 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 447.47 | 17.28 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.6 | 329.58 | 23.44 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 530.15 | 14.67 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)