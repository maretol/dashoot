package thread

import (
	"time"
)

// Thread はスレッド情報のコンストラクタです
type Thread struct {
	name   string  // スレッド名
	shoots []shoot // スレッドの書き込み
	// あと接続先とかも入れる必要がある？
}

type shoot struct {
	sentence string
	time     time.Time
}

// NewThread はコンストラクタです
func NewThread(name string) *Thread {
	thread := &Thread{name: name}
	return thread
}

//SomeInitialize は初期化処理です
func (thread *Thread) SomeInitialize() {

}

// GetThreadName はスレッド名を返すメソッド(いらない)
func (thread *Thread) GetThreadName() string {
	return thread.name
}

// PutNewShoots 新しい書き込み
func (thread *Thread) PutNewShoots(sentence string) {
	newshoot := shoot{sentence: sentence, time: time.Now()}
	if len(thread.shoots) > 50 {
		// 書き込みがすでに50件溜まっていた場合
		// themeThread.shoots から末尾を追い出してdatabaseCacheに渡してあげる
	}
	// TODO: 書き込みの処理。書き込んだあとにPUSHするところまで必要
	thread.checkSentence(newshoot.sentence)
}

// PushNewShoot は新しい書き込みをスレッド閲覧者に送り出すメソッドです。実際はmain/readパッケージ側がPUSHするので、PUSHする構造体とかを生成して通知する？仕組みになる
func (thread *Thread) PushNewShoot() {
	// 通知方法を考える。observableあたりが近いわけだけど
}

// PushDatabase は溜まった書き込みをDBに書き込むメソッド
func (thread *Thread) PushDatabase() {
	// これは並列処理で適宜後ろに回すのが良いと思われる
}

// GetFiftyShoots はまず最初にスレッドを開いたときに見る50件を取得するメソッド
func (thread *Thread) GetFiftyShoots() {
	// これは普通に themeThread.shoots でいいけど渡すときの構造体を決めてからね
}

// checkSentence は投稿された文章が不適切かどうかを判別します。まあチェッカーですよ
func (thread *Thread) checkSentence(sentence string) bool {
	return true
}
