package main

import (
	"encoding/json"
	"github.com/HiromasaNojima/explain-order-job-test/orderapi"
	"net/http"
)

func main() {
	// ...production環境用のコード...

	// ...development環境でのみデプロイされる、テスト用のコード...
	if env == Development {
		http.HandleFunc("/internal/order-job", func(w http.ResponseWriter, r *http.Request) {
			externalAPI := &orderapi.ExternalAPIMock{
				RequestOrderFunc: func(req orderapi.OrderInfo) error {
					return nil
				},
			}
			err := NewJobUsecase(
				// ...DI
				externalAPI,
			).Order()
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			// 発注用APIをリクエストした時の値を取得、レスポンスで返却し、受け取り側で意図した通りにリクエストされているか評価する。
			called := externalAPI.RequestOrderCalls()[0]
			b, err := json.Marshal(called)
			w.Write(b)
			return
		})
		http.ListenAndServe(":8080", nil)
	}
}

type jobUsecaseImpl struct {
	// ...dependency
	externalAPI orderapi.ExternalAPI
}

func (r *jobUsecaseImpl) Order() error {
	// ...発注ロジック
	req := orderapi.OrderInfo{
		ProductName: "発注したい商品の名前",
	}
	// ...発注APIの呼び出し
	err := r.externalAPI.RequestOrder(req)
	if err != nil {
		return err
	}
	return nil
}

func NewJobUsecase(externalAPI orderapi.ExternalAPI) JobUsecase {
	return &jobUsecaseImpl{
		// ...DI
		externalAPI: externalAPI,
	}
}

type JobUsecase interface {
	Order() error
}

var env = "development"

const Development = "development"
