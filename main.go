package main

import (
	"context"

	"solufit/go/internal/presenter"
)

// @title shiron notice API
// @version v0.1.0
// @description shiron system からユーザーチェインの紐づけ情報を使って各種リソース(webhock / Email / discord / slack)へ通知を送信するためのAPIです。
// @host localhost:8080
func main() {
	srv := presenter.NewServer()
	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}
}
