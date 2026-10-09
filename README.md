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
**Last Updated:** Fri, 09 Oct 2026 09:36:36 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 401.82 | 19.33 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 402.07 | 19.31 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 692.49 | 11.22 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.7.1 | 530.11 | 14.64 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 850.48 | 9.15 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)