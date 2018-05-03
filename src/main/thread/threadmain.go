package thread

import (
	"time"
)

type thread struct {
	name   string  // スレッド名
	shoots []shoot // スレッドの書き込み
	// あと接続先とかも入れる必要がある？
}

type shoot struct {
	sentence string
	time     time.Time
}

var themeThread thread
var databaseCache []shoot

// Init は初期化処理です。
func Init(name string) {
	themeThread = thread{name: name, shoots: make([]shoot, 50, 1000)}
}

// GetThreadName はスレッド名を返すメソッド
func GetThreadName() string {
	return themeThread.name
}

// PutNewShoots 新しい書き込み
func PutNewShoots(sentence string) {
	newshoot := shoot{sentence: sentence, time: time.Now()}
	if len(themeThread.shoots) > 50 {
		// 書き込みがすでに50件溜まっていた場合
		// themeThread.shoots から末尾を追い出してdatabaseCacheに渡してあげる
	}
	// TODO: 書き込みの処理。書き込んだあとにPUSHするところまで必要
	checkSentence(newshoot.sentence)
}

// PushNewShoot は新しい書き込みをスレッド閲覧者に送り出すメソッドです。実際はmain/readパッケージ側がPUSHするので、PUSHする構造体とかを生成して通知する？仕組みになる
func PushNewShoot() {
	// 通知方法を考える。observableあたりが近いわけだけど
}

// PushDatabase は溜まった書き込みをDBに書き込むメソッド
func PushDatabase() {
	// これは並列処理で適宜後ろに回すのが良いと思われる
}

// GetFiftyShoots はまず最初にスレッドを開いたときに見る50件を取得するメソッド
func GetFiftyShoots() {
	// これは普通に themeThread.shoots でいいけど渡すときの構造体を決めてからね
}

// checkSentence は投稿された文章が不適切かどうかを判別します。まあチェッカーですよ
func checkSentence(sentence string) bool {
	return true
}
