# ANTLR 工具

修改 `antlr/grammar/TideSQL.g4` 后需用 ANTLR 重新生成 Go parser。

```bash
make sql-gen        # 若缺少 jar 会先执行 make antlr-download
make antlr-download # 仅下载 antlr-4.13.2-complete.jar（约 2MB）
```

依赖：**JDK**（`java` 命令可用）。jar 来自 [antlr.org](https://www.antlr.org/download/)，已加入 `.gitignore`，不会提交到仓库。
