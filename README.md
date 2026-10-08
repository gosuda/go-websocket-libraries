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
**Last Updated:** Thu, 08 Oct 2026 09:26:39 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 344.45 | 22.53 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 434.16 | 17.89 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 695.12 | 11.18 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 471.27 | 16.47 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 875.49 | 8.90 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)