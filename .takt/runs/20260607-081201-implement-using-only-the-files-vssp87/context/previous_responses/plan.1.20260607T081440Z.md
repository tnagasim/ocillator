計画レポートを作成しました。以下が調査・設計の結果です。

---

## 計画サマリー

### 現状

| ファイル | 状態 |
|---------|------|
| `go.mod` | 存在・モジュール名 OK・go バージョンが `1.26.4`（要変更） |
| `.mise.toml` | 存在・`go = "1.26"`（要変更→`1.24`） |
| ディレクトリ構成 | `docs/` のみ。`cmd/`, `internal/`, `pkg/`, `configs/` は未作成 |

### 作成・変更するファイル（6種）

1. **`.mise.toml`** — `go = "1.26"` → `go = "1.24"` に変更
2. **`go.mod`** — `go 1.26.4` → `go 1.24` に変更
3. **`internal/.gitkeep`, `pkg/.gitkeep`, `configs/.gitkeep`** — 空ファイル新規作成
4. **`cmd/ocillator/main.go`** — cobra を使った CLI（ルート + `sync`/`release`/`deploy` プレースホルダ、`--version` は cobra の `Version` フィールドで実装）
5. **`.github/workflows/build.yml`** — test → cross-compile の 2-ジョブ構成
6. **`go.sum`** — `go get github.com/spf13/cobra@latest && go mod tidy` で自動生成

### 主な設計決定

- **cobra の `Version` フィールド**を使うことで `--version` フラグを自動生成（手動定義不要）
- **`cmd/ocillator/main.go` は単一ファイル**（タスク指定通り・推定 60〜80 行）
- **GitHub Actions は `test` → `build` の 2 ジョブ構成**（`build` は `needs: test`）・matrix でクロスコンパイル

### 注意点

- `go.sum` は手動作成せず `go mod tidy` で生成する
- `cmd/` には `.gitkeep` を置かない（タスク明示）
- GitHub Actions の `go-version` は `"1.24"`（文字列）で指定（数値 `1.24` は YAML で `1.2` に解釈される）