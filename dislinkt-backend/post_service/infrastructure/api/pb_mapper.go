package api

import (
	pb "github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/common/proto/post_service"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/post_service/domain"
)

func mapProduct(product *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:   product.Id.Hex(),
		Name: product.Name,
	}

	return postPb
}
