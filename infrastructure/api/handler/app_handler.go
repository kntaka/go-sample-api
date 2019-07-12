package handler

// バリデーションやミドルウエア処理などのレイヤー
type AppHandler interface {
	IndexHandler
	UserHandler
}

type appHandler struct {
	IndexHandler
	UserHandler
}

func NewAppHandler(ih IndexHandler,uh UserHandler) AppHandler {
	return &appHandler{ih, uh}
}
