# 技術スタック

## 言語・ランタイム
- **Go 1.21.1** - メインプログラミング言語
- **Windows専用** - Windowsプラットフォーム専用

## 依存関係
- `github.com/schollz/progressbar/v3` - プログレスバー表示
- `github.com/mattn/go-runewidth` - Unicode文字幅計算
- `github.com/mitchellh/colorstring` - ターミナル色彩サポート
- `github.com/rivo/uniseg` - Unicode文字分割
- `golang.org/x/sys` & `golang.org/x/term` - システム・ターミナルユーティリティ

## ビルドシステム
- **go-winres** - アイコンとマニフェスト用Windowsリソース埋め込みツール
- **go generate** - 自動リソース生成に使用

## 共通コマンド

### ビルド
```bash
# Windowsリソース生成
go generate

# 実行ファイルビルド
go build -o scoop-update-to-cleanup.exe

# 最適化付きビルド
go build -ldflags="-s -w" -o scoop-update-to-cleanup.exe
```

### 開発
```bash
# 直接実行
go run scoop-update-to-cleanup.go

# 依存関係インストール
go mod tidy

# 依存関係更新
go mod download
```

## プラットフォーム要件
- Windows 7以降
- Scoopパッケージマネージャーがインストール済み
- コマンドプロンプトアクセス