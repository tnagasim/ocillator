# OCIllator

## 実現すること

OCIllator は、インターネット接続可能な環境から取得した OCI イメージを、エアギャップ（オフライン）環境へ安全かつ効率的に配布するための CLI ツールです。

### 背景

現在の運用では以下のような構成を想定しています。

```text
GHCR
 ↓
Windows
 ↓
OCI Layout
 ↓
差分同期
 ↓
Offline VPC
 ↓
Deploy
```

### 解決したい課題

- オフライン環境へコンテナイメージを持ち込みたい
- イメージ更新時に差分のみ転送したい
- OCI標準形式を利用したい
- Dockerレジストリを必須にしたくない
- 配布・リリース作業を自動化したい
- 複数のイメージをまとめて管理したい

### 提供する機能（v1）

- GHCRなどのOCI Registryからイメージ取得
- OCI Layoutへの保存
- OCI Layoutの差分検出
- SFTPによる差分転送
- オフライン環境への配布
- リリース履歴管理
- YAMLによる設定管理

---

## 設計概要

### 全体構成

```text
                ┌──────────────┐
                │ ghcr.io      │
                └──────┬───────┘
                       │
                       ▼
              ┌─────────────────┐
              │ OCIllator       │
              │ Sync Engine     │
              └──────┬──────────┘
                     │
                     ▼
              ┌─────────────────┐
              │ OCI Layout      │
              │ Local Cache     │
              └──────┬──────────┘
                     │
               SFTP Sync
                     │
                     ▼
              ┌─────────────────┐
              │ Offline VPC     │
              │ OCI Layout      │
              └──────┬──────────┘
                     │
                     ▼
              ┌─────────────────┐
              │ Deploy Engine   │
              └─────────────────┘
```

### 技術方針

OCIllator 自身は OCI の実装を極力持たず、既存の成熟した OCI ツールを活用します。

```text
OCIllator
 ├─ Release Management
 ├─ Sync Management
 ├─ Manifest Management
 └─ Deployment Orchestration

regctl
 └─ Registry ↔ OCI Layout

skopeo
 └─ OCI Layout ↔ Runtime
```

### v1の責務

- イメージ一覧管理
- OCI Layout管理
- 差分判定
- 転送制御
- リリース管理
- デプロイ制御

### v1で実装しないもの

- OCI Registry実装
- OCI Manifest実装
- コンテナランタイム実装
- Kubernetesデプロイ機能

---

## 開発環境

### 言語

- Go

理由:

- 単一バイナリ配布
- クロスコンパイルが容易
- OCI関連ライブラリが充実
- CLIツールとの相性が良い

### 開発OS

- macOS

### ターゲットOS

- Windows
- Linux
- macOS

### バージョン管理

mise を利用して Go のバージョンを固定する。

`.mise.toml`

```toml
[tools]
go = "1.26.4"
```

### リポジトリ構成

```text
ocillator/
├── cmd/
├── internal/
├── pkg/
├── configs/
├── docs/
├── .mise.toml
├── go.mod
└── go.sum
```

### AIエージェント開発方針

Dev Container は利用しない。

理由:

- AIエージェントの並列実行がしづらい
- コンテナ起動コストがある
- Git Worktreeとの相性が悪い

### 推奨構成

```text
Mac
 ├─ mise
 ├─ direnv
 ├─ git worktree
 ├─ Go
 └─ AI Coding Agent
```

### Worktree構成

```text
ocillator/
├── main
├── feature-sync
├── feature-release
├── feature-deploy
└── feature-config
```

各 Worktree を独立した AI エージェントに割り当てる。

### CI/CD

GitHub Actions

生成物:

```text
ocillator-darwin-arm64
ocillator-linux-amd64
ocillator-windows-amd64.exe
```

### リリース形式

単一実行ファイル

```text
Windows
 └─ ocillator.exe

Linux
 └─ ocillator

macOS
 └─ ocillator
```

利用者は Go や追加ランタイムをインストールする必要がない。