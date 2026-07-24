テスト作成が完了しました。

## 作業サマリー

### 作成したファイル
- **`cmd/ocillator/main_test.go`** — cobra CLI の振る舞いテスト（6ケース）

### テストの設計判断

| 判断 | 理由 |
|------|------|
| `newRootCmd()` を前提とする | package main のテストで各ケースが独立した状態を持てる。draft ステップの実装インターフェースを明確に規定できる |
| `*cobra.Command` に `SetOut(&buf)` でキャプチャ | cobra の version/help 出力は `OutOrStdout()` 経由のため `SetOut` で確実にキャプチャできる |
| 未知コマンドテストで `SilenceErrors/Usage` | テスト出力にエラーメッセージが混入するのを防ぐ |
| インテグレーションテスト不要 | 横断データフロー・状態合流・オプション伝搬がいずれも存在しない |

### テスト実行結果
`undefined: newRootCmd` のみ（`main.go` 未実装による想定内エラー）。import パスミスなど実装後も残る問題はなし。