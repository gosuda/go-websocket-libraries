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
**Last Updated:** Tue, 15 Jul 2025 03:50:31 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 205.22 | 37.68 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 293.73 | 26.34 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 453.46 | 17.02 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.4 | 357.26 | 21.66 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 547.77 | 14.18 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)