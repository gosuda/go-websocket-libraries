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
**Last Updated:** Tue, 20 Oct 2026 07:55:27 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 245.54 | 31.52 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 286.18 | 27.06 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 493.90 | 15.65 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 343.40 | 22.54 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 584.34 | 13.31 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)