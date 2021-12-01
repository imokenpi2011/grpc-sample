# grpc-sample
参考資料 : https://github.com/ishishow/grpc_tutorial
- [grpc-sample](#grpc-sample)
  - [gRPCとは？](#grpcとは)
    - [ざっくり](#ざっくり)
    - [特徴](#特徴)
  - [疑問点](#疑問点)
  - [参考記事(自分用)](#参考記事自分用)
  - [参考ソース](#参考ソース)

## gRPCとは？
### ざっくり
HTTP/2を使用しており、データがバイナリ形式のままやり取りされるため、軽量かつ通信量を抑えることができる。

高トラフィックのソシャゲ開発などにも使われているみたい。
### 特徴
基本的にはGoogleのProtocol Buffersを使用して通信を行う。

基本的に.protoファイルに記載された内容からGolangのファイルを自動生成し、それを呼び出して使用する。

使い方のイメージ的には
1. 実装した関数を定義したサービスをprotoファイルに定義
2. `protoc --go_out=[出力フォルダ] [各オプション] [protoファイル]`でソースを自動生成
3. server側のソースにサービスを登録して起動
4. client側のソースにサービスを登録し、サーバー側とやり取りする

.protoファイルにデータ定義や関数をあらかじめ定義するので、どういったAPIを使用しているかがわかりやすい


## 疑問点
- protoファイルを使った自動生成が安定しない。goのパス設定とうまく噛み合ってない
- サーバー側とクライアント側にそれぞれmainパッケージのmainメソッドを定義しているためapp直下に置かない方がいい？
- 巨大なプロジェクトにおいて使用する場合のイメージがまだ掴めてない。普通にREST APIで開発して慣れた方がいいかも

## 参考記事(自分用)
- https://grpc.io/docs/languages/go/quickstart/
- https://qiita.com/ishishow/items/c9b633d3d4d22db74a6d
- https://zenn.dev/mikankitten/articles/0437fa6fb7de82

## 参考ソース
- https://github.com/grpc-ecosystem/grpc-gateway
- https://github.com/eddycjy/grpc-hello-world