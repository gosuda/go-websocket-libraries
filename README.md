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
**Last Updated:** Sun, 05 Oct 2025 03:27:40 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.14 | 215.48 | 35.95 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 277.10 | 27.93 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 440.09 | 17.55 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.7 | 334.74 | 23.16 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.8.9 | 542.07 | 14.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)