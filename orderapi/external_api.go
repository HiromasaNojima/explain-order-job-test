package orderapi

type ExternalAPI interface {
	RequestOrder(request OrderInfo) error
}
type externalAPIImpl struct{}

type OrderInfo struct {
	// ..発注情報
	ProductName string `json:"product_name"`
}

func (r externalAPIImpl) RequestOrder(request OrderInfo) error {
	// 外部API呼び出し
	return nil
}
