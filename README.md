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
**Last Updated:** Sat, 24 Oct 2026 07:35:41 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 320.12 | 24.19 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 346.99 | 22.38 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 640.43 | 12.11 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 425.90 | 18.25 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 759.54 | 10.25 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)