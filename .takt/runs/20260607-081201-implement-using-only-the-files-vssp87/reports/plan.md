# タスク計画

## 元の要求

`.takt/runs/20260607-081201-implement-using-only-the-files-vssp87/context/task/order.md` に基づき、OCIllator プロジェクトの初期ディレクトリ構成・Go モジュール・CLI エントリポイント・CI/CD を構築する。参照資料: `docs/project-proposal.md`。

## 分析結果

### 目的

Go CLI ツール「OCIllator」のプロジェクトスキャフォールドを作成する。具体的には、Go モジュール設定・ディレクトリ骨格・cobra を使った CLI エントリポイント・GitHub Actions CI/CD を整備し、`go build ./...` と `go run ./cmd/ocillator --help/--version` が通る状態にする。

### 分解した要件

| # | 要件 | 種別 | 備考 |
|---|------|------|------|
| 1 | `.mise.toml` の Go バージョンを `1.24` に固定する | 明示 | 現状 `go = "1.26"` → `go = "1.24"` |
| 2 | `go.mod` の go ディレクティブを `1.24` にする | 暗黙 | 明示要求 #1「Go 1.24 を固定する」から導出。現状 `go 1.26.4` |
| 3 | `internal/` ディレクトリを `.gitkeep` 付きで作成する | 明示 | |
| 4 | `pkg/` ディレクトリを `.gitkeep` 付きで作成する | 明示 | |
| 5 | `configs/` ディレクトリを `.gitkeep` 付きで作成する | 明示 | |
| 6 | `cmd/` ディレクトリを作成する（`.gitkeep` なし） | 明示 | タスク明示「`cmd/` を除く」 |
| 7 | `cmd/ocillator/main.go` にルートコマンドを定義する（cobra 使用） | 明示 | |
| 8 | `ocillator --version` で `v0.0.1` を出力する | 明示 | cobra の `Version` フィールドで実装 |
| 9 | `sync` サブコマンドを空実装で登録する | 明示 | |
| 10 | `release` サブコマンドを空実装で登録する | 明示 | |
| 11 | `deploy` サブコマンドを空実装で登録する | 明示 | |
| 12 | cobra を `go.mod` の依存に追加し `go.sum` を生成する | 暗黙 | 明示要求 #7「cobra 推奨でCLI実装」から導出 |
| 13 | `.github/workflows/build.yml` を作成する（`push`/`pull_request` トリガー） | 明示 | |
| 14 | CI で `go test ./...` をビルド前に実行する | 明示 | |
| 15 | CI で `darwin/arm64` 向けに `ocillator-darwin-arm64` をビルドする | 明示 | |
| 16 | CI で `linux/amd64` 向けに `ocillator-linux-amd64` をビルドする | 明示 | |
| 17 | CI で `windows/amd64` 向けに `ocillator-windows-amd64.exe` をビルドする | 明示 | |
| 18 | CI の Go バージョンを `1.24` に固定する | 明示 | |

### 参照資料の調査結果

`docs/project-proposal.md` を確認。以下の点が現在の実装と差異あり：

- **Go バージョン:** `project-proposal.md` は `go = "1.24"` を指定。現状の `.mise.toml` は `1.26`、`go.mod` は `1.26.4`
- **ディレクトリ構成:** `project-proposal.md` が示す `cmd/`, `internal/`, `pkg/`, `configs/` は未作成
- **CLI・CI/CD:** 未実装

`project-proposal.md` はスキャフォールドの設計方針（言語 Go・cobra CLI・GitHub Actions・クロスコンパイル成果物）を規定しており、`order.md` の各タスクはその仕様を実装に落とす指示として整合している。

### スコープ

| ファイル | 変更種別 |
|---------|---------|
| `.mise.toml` | 変更（go バージョン） |
| `go.mod` | 変更（go バージョン） |
| `go.sum` | 新規（`go mod tidy` で生成） |
| `internal/.gitkeep` | 新規 |
| `pkg/.gitkeep` | 新規 |
| `configs/.gitkeep` | 新規 |
| `cmd/ocillator/main.go` | 新規 |
| `.github/workflows/build.yml` | 新規 |

既存の `docs/`, `LICENSE`, `.gitignore` は変更しない。

### 検討したアプローチ（設計判断がある場合）

| アプローチ | 採否 | 理由 |
|-----------|------|------|
| `cmd/ocillator/main.go` 単一ファイルに CLI 全体を収める | **採用** | タスクの対象ファイルが `main.go` 1 つのみ。サブコマンドは空実装のみで ~70 行以内に収まる |
| `cmd/ocillator/root.go` を分離する | 不採用 | 現時点では過剰分割。サブコマンドに実装が入る段階で分離すればよい |
| cobra の `Version` フィールドで `--version` を実装 | **採用** | cobra が `--version` フラグを自動生成。手動フラグ定義不要でシンプル |
| GitHub Actions: test と build を単一ジョブにする | 不採用 | `go test` 失敗時もビルドが走ってしまう。test → build の 2 ジョブ構成が明確 |
| GitHub Actions: build を matrix で並列実行 | **採用** | 3 プラットフォームを宣言的に記述でき、成果物ごとに artifact を分離しやすい |

### 実装アプローチ

1. `.mise.toml` の `go = "1.26"` を `go = "1.24"` に変更する
2. `go.mod` の `go 1.26.4` を `go 1.24` に変更する
3. `internal/.gitkeep`、`pkg/.gitkeep`、`configs/.gitkeep` を空ファイルとして作成する
4. `go get github.com/spf13/cobra@latest && go mod tidy` を実行し、`go.mod` に cobra を追加、`go.sum` を生成する
5. `cmd/ocillator/main.go` を実装する（cobra ルートコマンド + 3 サブコマンド空実装）
6. `.github/workflows/build.yml` を作成する（test ジョブ → build ジョブ matrix 構成）
7. `go build ./...` と `go run ./cmd/ocillator --help`、`go run ./cmd/ocillator --version` で動作確認する

### 到達経路・起動条件

| 項目 | 内容 |
|------|------|
| 利用者が到達する入口 | シェルから `ocillator <subcommand>` を直接実行。GitHub Actions は `push` / `pull_request` イベントでトリガー |
| 更新が必要な呼び出し元・配線 | なし（新規 CLI。既存コードへの配線変更なし） |
| 起動条件 | `ocillator` バイナリが PATH に存在すること、または `go run ./cmd/ocillator` で実行 |
| 未対応項目 | なし |

## 実装ガイドライン

### `cmd/ocillator/main.go`

- cobra の `rootCmd.Version = "v0.0.1"` を設定することで `--version` フラグが自動生成される。`cobra.Command` に `Version` フィールドをセットする以外の実装は不要
- サブコマンドは `RunE` で `return nil` のみの空実装にする（`Run` ではなく `RunE` を使うことでエラーハンドリングの拡張が容易）
- `main()` では `rootCmd.Execute()` の返却エラーを受け取り、エラー時は `os.Exit(1)` する。cobra はデフォルトで `Execute()` 内部でエラーを stderr に出力するため、`main()` 側での追加出力は不要

### `go.mod` / `go.sum`

- `go.sum` は手動作成しない。必ず `go mod tidy` で生成する
- `go mod tidy` 実行後、`go.mod` の go ディレクティブが cobra の要求する最小バージョンに引き上げられる可能性がある。その場合は引き上げ後のバージョンを採用する（cobra v1.x は Go 1.15+ 対応のため 1.24 が下回ることはない）

### `.github/workflows/build.yml`

- `go-version:` の値は必ず文字列 `"1.24"` で記述する。YAML の数値リテラル `1.24` は `1.2` と解釈されるため
- `actions/checkout`・`actions/setup-go`・`actions/upload-artifact` はいずれも v4 を使用する
- build ジョブに `needs: test` を設定し、テスト失敗時はビルドをスキップする
- 成果物のバイナリ名は matrix の `artifact` 変数から参照し、`-o` オプションで指定する

### アンチパターン

- `cmd/` に `.gitkeep` を置かない（タスク明示事項）
- `--version` フラグを手動定義しない（cobra の自動生成を使う）
- `go.sum` を手動作成・編集しない

## スコープ外

| 項目 | 除外理由 |
|------|---------|
| OCI Registry 実装 | タスク明示「やらないこと」 |
| OCI Manifest 実装 | タスク明示「やらないこと」 |
| コンテナランタイム実装 | タスク明示「やらないこと」 |
| Kubernetes デプロイ機能 | タスク明示「やらないこと」 |
| Dev Container 設定 | タスク明示「やらないこと」 |
| `sync`/`release`/`deploy` サブコマンドの実際の実装 | タスク明示「実装は空でよい」 |

## 確認事項

- `go mod tidy` 後に go ディレクティブが `1.24` より上に引き上げられた場合、そのバージョンを採用してよいか（cobra v1.x は Go 1.15+ 対応のため実際には引き上げは発生しない見込み）