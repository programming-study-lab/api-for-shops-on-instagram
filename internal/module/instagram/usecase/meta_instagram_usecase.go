package usecase

import (
	"context"
)

type metaInstagramUsecase struct {
	// metaInstagramHttp *http.MetaInstagramHandler
	metaInstagramHttp interface{}
}

func NewMetaInstagramUsecase(metaInstagramHttp interface{}) *metaInstagramUsecase {
	return &metaInstagramUsecase{
		metaInstagramHttp: metaInstagramHttp,
	}
}

func (metaIG *metaInstagramUsecase) MediaUsecase(ctx context.Context) (map[string]any, error) {
	// metaIG.metaInstagramHttp.Media()

	return nil, nil
}
