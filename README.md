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
**Last Updated:** Mon, 05 Oct 2026 09:35:49 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 338.55 | 22.90 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 371.32 | 20.90 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 651.57 | 11.91 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.0 | 450.12 | 17.27 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 782.42 | 9.95 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)