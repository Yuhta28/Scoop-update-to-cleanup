# Scoop-update-to-cleanup
Go app that performs everything from updating Scoop packages to deleting cache in one go.

## セキュリティについて

### ファイルの検証
ダウンロードしたファイルの整合性を確認するため、以下のコマンドでハッシュ値を確認してください：

```cmd
certutil -hashfile scoop-update-to-cleanup.exe SHA256
```

最新リリースのハッシュ値は[Releases](https://github.com/[your-username]/scoop-update-to-cleanup/releases)ページで確認できます。

### Windows Defenderの警告について
このソフトウェアは署名されていないため、Windows Defenderで警告が表示される場合があります。
「詳細情報」→「実行」で実行可能です。全てのソースコードは公開されており、検証可能です。

## references 
https://zenn.dev/yuta28/articles/windows-scoopupdate-go-lang
