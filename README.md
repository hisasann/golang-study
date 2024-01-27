# golang-study

## Editors

### GoLand

あまりカスタムせず以下のプラグインをインストールした。

- IdeaVim
- Fig
- Copilot
- Material Theme UI

Go Modules を有効にする。

Settings -> Go -> Go Modules -> Enable Go Modules integration にチェックを入れる。

これで、 `logger.trace` のような他の pkg のメソッド定義位置に飛べなかった部分が飛べるようになる。

### VSCode

[Goland から VSCode への移行](https://zenn.dev/tellernovel_inc/articles/8a1ac1f1652605)

を見ながら、 VSCode 上のコマンドパレットを使って必要な拡張をインストールした。

**Go: Install/Update tools**

```shell
Tools environment: GOPATH=/Users/hisasann/go
Installing 7 tools at /Users/hisasann/go/bin
  gopls
  gotests
  gomodifytags
  impl
  goplay
  dlv
  staticcheck

Installing golang.org/x/tools/gopls@latest (/Users/hisasann/go/bin/gopls) SUCCEEDED
Installing github.com/cweill/gotests/gotests@v1.6.0 (/Users/hisasann/go/bin/gotests) SUCCEEDED
Installing github.com/fatih/gomodifytags@v1.16.0 (/Users/hisasann/go/bin/gomodifytags) SUCCEEDED
Installing github.com/josharian/impl@v1.1.0 (/Users/hisasann/go/bin/impl) SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay@v1.0.0 (/Users/hisasann/go/bin/goplay) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv@latest (/Users/hisasann/go/bin/dlv) SUCCEEDED
Installing honnef.co/go/tools/cmd/staticcheck@latest (/Users/hisasann/go/bin/staticcheck) SUCCEEDED

All tools successfully installed. You are ready to Go. :)
```

gofumpt は settings.json に以下を追加してみた。

```
"go.useLanguageServer": true,
"gopls": {
    "formatting.gofumpt": true
}
```
