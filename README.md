# 浜松市 新型コロナウイルス感染症対策サイト

## ライセンス
本ソフトウェアは、[MITライセンス](./LICENSE.txt)の元提供されています。


## コミュニケーションへの参加方法

浜松版Covid19対策サイトの運営方法や開発に関する議論はSlackにて行っています。

Slack アカウントを持っていない場合、[こちら](https://join.slack.com/t/jaws-ug-hamamatsu/shared_invite/zt-dc5cgs87-cNw8QOxXeqhJnB8YB8A6Uw)から登録してください.


`#code_for_hamamatsu` チャンネルにご参加ください。


## 開発者向け情報

### 環境構築の手順

- 必要となるNode.jsのバージョン: 10.19.0以上

**yarn を使う場合**
```bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev
```

**docker compose を使う場合**
```bash
# serve with hot reload at localhost:3000
$ docker-compose up --build
$ docker exec -it covid19 sh

# コンテナに入ってから
/app yarn install
/app yarn dev
```

**Vagrant を使う場合**
```bash
# serve with hot reload at localhost:3000
$ vagrant up
```

### `Cannot find module ****` と怒られた時

**yarn を使う場合**
```bash
$ yarn install
```

**docker compose を使う場合**
```bash
$ docker-compose run --rm app yarn install
```

起動後、[http://localhost:3000](http://localhost:3000) へアクセスする。

### VSCode + Remote Containersで開発する場合

1. VSCodeの拡張機能「[Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)」を導入します。
2. [この画像（外部サイト）](https://code.visualstudio.com/docs/remote/containers#_quick-start-try-a-dev-container)のように左下部の「Open Folder in Container」でこのリポジトリのルートを選択すれば環境構築が始まります。

### E2Eテスト

```
yarn test:e2e
```

#### Topic
- 設定を変更したい場合は、`.devcontainer/devcontainer.json`を修正してください。
詳細は[devcontainer.jsonのリファレンス](https://code.visualstudio.com/docs/remote/containers#_devcontainerjson-reference)を参照してください。
- Remote Container実行時のみ有効な拡張機能「ESLint」を導入していますが、必要に応じて`devcontainer.json`の`extensions`に追加してください。
詳細な手順は[こちら（外部サイト）](https://code.visualstudio.com/docs/remote/containers#_managing-extensions)を参照してください。
- 開発環境を再構築する場合は、左下部の「Rebuild Container」を実行してください。

### 本番環境/その他の判定

`process.env.GENERATE_ENV` の値が、本番の場合は`'production'`に、それ以外の場合は `'development'` になっています。  
テスト環境のみで実行したい処理がある場合はこちらの値をご利用ください。

### 本番環境への反映

main ブランチがアップデートされると、自動的にNetlifyにより、本番サイト https://stopcovid19.code4hamamatsu.org/ が更新されます。

### ブランチルール

main 以外は Pull Request は禁止です。
Pull Request を送る際の branch は、以下のネーミングルールでお願いします。

機能追加系： feature/#{ISSUE_ID}-#{branch_title_name}  
ホットフィックス系: hotfix/#{ISSUE_ID}-#{branch_title_name}

#### 基本的なブランチ

**2020/06/06**

新機能開発のペースが落ち着いてきたため、devブランチを一旦廃止します。

今後のプルリクエストは main 宛に送ってください。

main へプルリクエストを送るとNetlifyによって検証用のURLが発行されるため、そのURLから変更内容の確認ができます。

| 目的 | ブランチ | 確認URL | 備考 |
| ---- | -------- | ---- | ---- |
| 開発(2020/06/06廃止) | dev-hamamatsu | https://dev-stopcovid19.code4hamamatsu.org/ | base branch。~~基本はこちらに Pull Requestを送ってください~~ |
| 本番 | main | https://stopcovid19.code4hamamatsu.org/ | base branch こちらからブランチ切って Pull Requestを送ってください |

#### システムで利用しているブランチ
| 目的 | ブランチ | 確認URL | 備考 |
| ---- | -------- | ---- | ---- |
| 本番サイトHTML | main | https://stopcovid19.code4hamamatsu.org/ | 静的ビルドされたHTMLが置いてある場所 |
