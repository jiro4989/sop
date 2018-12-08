# sop (safety operation)

安全なオペレーションをしたい。

## 目的

AnsibleやDockerを導入するほどではないけれど、でもやっぱり忘れずバックアップはと
っておいて、万が一に備えておきたい、といった僕の悩みを解決するために作成しました
。

## インストール

GitHub Releaseページからダウンロードするか、Goをインストールしている方は以下のコ
マンドを実行してインストールします。

```bash
go get github.com/jiro4989/sop
```

## 使い方

簡単なコピーの例

```bash
sop cp src.txt dst.txt
sop cp src.txt dst.txt
sop cp src.txt dst.txt
```

実行結果の確認

```bash
ls -la *.txt*
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 dst.txt
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 dst.txt.2018-12-08_151322
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 dst.txt.2018-12-08_151339
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 src.txt
```

見ての通り、copyによって上書きされたdst.txtのバックアップが保存されています。
やりたいことはこれだけです。

copyの際にユーザ、グループ、権限を指定できます。

```bash
sudo sop cp src.txt dst.txt -o root -g syslog -m 0740
```

## やりたいこと

cpやvimで編集するときに、自動でバックアップをとりたい。
差分を比較して、差分があれば更新したい。
