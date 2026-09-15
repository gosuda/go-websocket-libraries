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
**Last Updated:** Tue, 15 Sep 2026 08:01:45 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 333.11 | 23.32 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 358.36 | 21.66 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 632.75 | 12.26 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 440.05 | 17.64 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 761.26 | 10.23 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)