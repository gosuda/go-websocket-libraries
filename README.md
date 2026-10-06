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
**Last Updated:** Tue, 06 Oct 2026 09:24:04 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 386.11 | 20.13 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 409.71 | 18.98 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 696.91 | 11.12 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 527.58 | 14.74 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 869.92 | 8.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)