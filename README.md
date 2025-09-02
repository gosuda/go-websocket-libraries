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
**Last Updated:** Tue, 02 Sep 2025 03:29:50 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.13 | 188.82 | 40.94 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 272.91 | 28.35 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 431.64 | 17.90 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 335.20 | 23.11 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 523.26 | 14.85 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)