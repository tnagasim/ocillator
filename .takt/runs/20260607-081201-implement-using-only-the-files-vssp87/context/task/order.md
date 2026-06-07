## タスク指示書：OCIllator プロジェクト初期スキャフォールド

### 概要

`docs/project-proposal.md` に基づき、OCIllator プロジェクトの初期ディレクトリ構成・Go モジュール・CLI エントリポイント・CI/CD を構築する。

---

### 参照資料

- `docs/project-proposal.md`

---

### タスク一覧

#### 【優先度：高】Go モジュール・ディレクトリ初期化

**対象ファイル:**
- `go.mod`
- `go.sum`
- `.mise.toml`

**作業内容:**
- `go mod init` でモジュールを初期化する（モジュール名は `github.com/tnagasim/ocillator` を想定）
- `.mise.toml` に Go `1.24` を固定する

```toml
[tools]
go = "1.24"
```

**確認方法:**
- `go build ./...` がエラーなく通ること

---

#### 【優先度：高】ディレクトリ構成の作成

**対象ディレクトリ:**
```
cmd/
internal/
pkg/
configs/
docs/   ← 既存
```

**作業内容:**
- 上記ディレクトリを作成する
- 各ディレクトリに `.gitkeep` を配置して空ディレクトリをコミット可能にする（`cmd/` を除く）

---

#### 【優先度：高】CLI エントリポイントの実装

**対象ファイル:**
- `cmd/ocillator/main.go`

**作業内容:**
- `main.go` を作成し、CLI フレームワーク（`cobra` 推奨）でルートコマンドを定義する
- `ocillator --version` で版番号（`v0.0.1` 等）を出力できること
- サブコマンドは今後追加予定のプレースホルダとして `sync`・`release`・`deploy` を登録する（実装は空でよい）

**確認方法:**
- `go run ./cmd/ocillator --help` でヘルプが表示されること
- `go run ./cmd/ocillator --version` でバージョンが表示されること

---

#### 【優先度：中】GitHub Actions CI/CD の設定

**対象ファイル:**
- `.github/workflows/build.yml`

**作業内容:**
- `push` / `pull_request` トリガーで以下の成果物をクロスコンパイルするワークフローを作成する

| OS | アーキテクチャ | 成果物 |
|----|---------------|--------|
| darwin | arm64 | `ocillator-darwin-arm64` |
| linux | amd64 | `ocillator-linux-amd64` |
| windows | amd64 | `ocillator-windows-amd64.exe` |

- Go バージョンは `1.24` に固定する
- `go test ./...` をビルド前に実行する

**確認方法:**
- ワークフローファイルの YAML 構文が valid であること（`actionlint` または `yamllint` で確認）

---

### やらないこと（ユーザー明示）

- OCI Registry 実装
- OCI Manifest 実装
- コンテナランタイム実装
- Kubernetes デプロイ機能
- Dev Container の利用

---

### Open Questions

- モジュール名（`github.com/tnagasim/ocillator`）は正しいか。GitHub リポジトリ URL が確定していれば教えてください。