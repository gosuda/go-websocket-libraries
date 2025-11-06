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
**Last Updated:** Thu, 06 Nov 2025 03:38:26 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 212.19 | 36.54 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.27 | 27.85 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 436.65 | 17.72 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 335.39 | 23.08 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 529.26 | 14.70 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)