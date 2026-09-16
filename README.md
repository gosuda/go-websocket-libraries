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
**Last Updated:** Wed, 16 Sep 2026 07:54:32 UTC

| Library                                         | Version         | Throughput (MB/s) | Avg Latency (ms) |
| ----------------------------------------------- | --------------- | ----------------- | ---------------- |
| [coder/websocket](https://github.com/coder/websocket) | v1.8.15 | 255.84 | 30.18 |
| [gobwas/ws](https://github.com/gobwas/ws) | v1.4.0 | 282.21 | 27.42 |
| [gorilla/websocket](https://github.com/gorilla/websocket) | v1.5.3 | 495.05 | 15.65 |
| [lesismal/nbio](https://github.com/lesismal/nbio) | v1.6.12 | 347.04 | 22.34 |
| [lxzan/gws](https://github.com/lxzan/gws) | v1.10.2 | 582.78 | 13.35 |
<!-- BENCHMARK_TABLE_END -->

**Performance Over Time: Throughput (MB/s):**

![Benchmark Performance Graph](benchmark_performance.png)