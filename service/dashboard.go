package service

import (
	"context"

	pb "auth-athlevo/genproto/auth"
	st "auth-athlevo/internal/storage/postgres"
)

type DashboardService struct {
	storage st.Storage
	pb.UnimplementedDashboardServiceServer
}

func NewDashboardService(storage *st.Storage) *DashboardService {
	return &DashboardService{
		storage: *storage,
	}
}

func (s *DashboardService) GetPersonalAccessCount(ctx context.Context, req *pb.AccessCountReq) (*pb.AccessCountRes, error) {
	res, err := s.storage.DashboardS.GetPersonalAccessCount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s *DashboardService) GetTotalPersonalBookingRevenue(ctx context.Context, req *pb.TotalRevenueReq) (*pb.TotalRevenueRes, error) {
	res, err := s.storage.DashboardS.GetTotalPersonalBookingRevenue(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *DashboardService) GetAccessCountBySubscriptionID(ctx context.Context, req *pb.SubscriptionCountReq) (*pb.SubscriptionCountRes, error) {
	res, err := s.storage.DashboardS.GetAccessCountBySubscriptionID(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *DashboardService) GetBookingRevenueBySubscriptionID(ctx context.Context, req *pb.TotalRevenueBySubscriptionReq) (*pb.TotalRevenueBySubscriptionRes, error) {
	res, err := s.storage.DashboardS.GetBookingRevenueBySubscriptionID(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}
