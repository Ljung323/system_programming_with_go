
# 14 Go 言語と並列処理

## 14.1 複数の仕事を同時に行うとは?

- 並行（concurrent）: CPU数、コア数の限界を超えて複数の仕事を同時に行う（スループットは変わらない）
- 並列（parallel）: 複数のCPU、コアを効率よく扱って計算速度を上げる（スループットが向上する）

>Concurrency is the ability of two or more threads to execute in overlapping time periods. 
Parallelism is the ability to execute two or more threads simultaneously.
(訳)並行処理は、複数個のスレッドを共通の期間内で実行する能力のことです。
並列処理は、複数個のスレッドを同時に実行する能力のことです。
[出典:書籍 Linux System Programming, 2nd Edition Chap.7](https://learning.oreilly.com/library/view/linux-system-programming/9781449341527/)

[並行処理と並列処理](https://zenn.dev/hsaki/books/golang-concurrency/viewer/term)

- CPU における処理時間が大きい場合(ユーザー時間が支配的な場合)は並列
- I/O 待ちなどで CPU が暇をしているときは並行
で処理するというのが基本

>シングルコアで並行処理をする場合、トータルでのスループットは変わりません。

<img src="./スクリーンショット 2023-09-18 21.35.30.png">

## 14.2 Go言語の並列処理のための道具

goではgoroutine、チャネルを使って並列処理を簡単に書ける

- goroutine: [goのランタイムによって管理される軽量なスレッド](https://go.dev/tour/concurrency/1)
- チャネル: キューに並列処理用の「並列でアクセスされても正しく処理される」ことを保証する機能を組み合わせたもの(4章 p62)
   - チャネルは、データを順序よく受け渡すためのデータ構造である
   - チャネルは、並列処理されても正しくデータを受け渡す同期機構である
   - チャネルは、読み込み・書き込みで準備ができるまでブロックする機能である

-> Go言語では、goroutine間の情報共有方法としてチャネルを使うことを推奨

余談: [goのドキュメント](https://go.dev/doc/effective_go#concurrency)ではgoroutineもchannelもconcurrencyの項目で説明されている

### 14.2.1 goroutineと情報共有

goroutineを実行する親スレッドと子の間でデータのやり取り

親→子（goroutine）のデータ送信

- 関数の引数として渡す
- クロージャのローカル変数にキャプチャして渡す
   - 内部的には、無名関数に暗黙の引数が追加され、その暗黙の引数にデータや参照(変数は参照扱い)が渡されてgoroutineとして扱われる
   - 基本的には「引数で渡す」のと同じ
   - forループ内でgoroutineを起動する場合は「引数で渡す」べき
      - 単純なループに比べてgoroutineの起動が遅いため、別のデータが参照されることがあるため

子（goroutine）→親のデータ送信

- 引数やクロージャで渡したデータ構造に書き込む
   - 配列やマップ、チャネルなど
- クロージャでキャプチャした変数に書き込む
   - キャプチャはポインタを引数に渡した扱いになる


## 14.3 スレッドとgoroutineの違い

- スレッド
   - CPUコアに対してマッピングされる
   - 重いが高機能

- goroutine
   - OSのスレッド(Go製のアプリケーションから見ると1つの仮想CPU)にマッピングされる
   - 軽いが低機能
      - 起動時間、切り替え時間など、どれをとってもOSのスレッドの1000倍のオーダーで高速

## 14.4 GoのランタイムはミニOS

既にOS側でスレッドという機能が提供されている上で、goroutineという機能を持つメリット

- 大量のクライアントを効率よくさばくサーバーを実装する(いわゆるC10K)ときに、クライアントごとに1つのgoroutineを割り当てるような実装であっても、リーズナブルなメモリ使用量で処理できる
   - [C10K問題](https://ja.wikipedia.org/wiki/C10K%E5%95%8F%E9%A1%8C): Apache等クライアントの接続1つ1つにプロセスを割り当てるWebサーバだと、大量に接続が来た際に同時起動可能プロセス数を超えてしまい、新たに接続ができなくなる
- OSのスレッドでブロッキングを行う操作をすると、他のスレッドが処理を開始するにはOSがコンテキストスイッチをして順番を待つ必要があるが、Goの場合はチャネルなどでブロックしたら、残ったタイムスライスでランキューに入った別のgoroutineのタスクも実行できる
- プログラムのランタイムが、プログラム全体の中でどのgoroutineがブロック中なのかといった情報をきちんと把握しているため、デッドロックを作ってもランタイムが検知してどこでブロックしているのかを一覧で表示できる

## 14.5 runtimeパッケージのgoroutine関連の機能

runtimeパッケージ: goroutineを使った、より低レベルの操作を可能にする

### 14.5.1 runtime.LockOSThread()/runtime.UnlockOSThread()

現在実行中のOSスレッドでのみgoroutineが実行されるように束縛

- メインスレッドでの実行強制等
- Goのランタイムでは、シグナルを受け取るスレッドを固定するために使用

### 14.5.2 runtime.Gosched()

現在実行中のgoroutineを一時中断して、他のgoroutineに処理を回す

### 14.5.3 runtime.GOMAXPROCS(n)/runtime.NumCPU()

同時に実行するOSスレッド数(I/Oのブロック中のスレッドは除く)を制御する

## 14.6 RaceDetector


## 14.7 syncパッケージ

## 14.8 sync/atomicパッケージ

## 14.9 本章のまとめと次章予告
