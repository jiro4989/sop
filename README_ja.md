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

## できること

- backup
- cp
- edit {editor} {targetFile}
- rm

ヘルプの結果が以下の通り。

    sop is safety operation

    Usage:
      sop [flags]
      sop [command]

    Available Commands:
      backup      backup copy file as backup.
      cp          cp copies file and save backup file.
      edit        edit can edit file with your favorite editor, and save backup.
      help        Help about any command
      rm          rm removes file and save backup.

    Flags:
      -h, --help   help for sop

    Use "sop [command] --help" for more information about a command.


## 使い方

### 簡単なコピーの例

以下のようにコマンドを実行する。

```bash
sop cp src.txt dst.txt
sop cp src.txt dst.txt
sop cp src.txt dst.txt
```

実行結果の確認

```
ls -la *.txt*
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 dst.txt
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 dst.txt.2018-12-08_151322
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 dst.txt.2018-12-08_151339
-rw-rw-r-- 1 jiro jiro 5 12月  8 15:13 src.txt
```

見ての通り、copyによって上書きされたdst.txtのバックアップが保存されています。
やりたいことはこれだけです。

### ユーザ、グループ、権限指定の例

権限や所有者を指定しつつコピー。
Linuxコマンドでいう`install`コマンドと同じようなもの。

```bash
sudo sop cp src.txt dst.txt -o root -g syslog -m 0740
```

結果確認

```
ls -la *.txt
-rwxr-x-w- 1 root syslog 2 12月  8 15:44 dst.txt
-rwxr-x-w- 1 root syslog 2 12月  8 15:44 dst.txt.2018-12-08_154442
-rw-rw-r-- 1 jiro jiro   2 12月  8 15:44 src.txt
```

