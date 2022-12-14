package presenter

import (
	pb "product_service_composite/service"
)

func AdaptGrpcToEntityInt64(count *pb.DeleteResponse) int64 {
	return count.Count
}
