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

func (s *DashboardService) GetDailyPersonalAccessCount(ctx context.Context, req *pb.AccessCountReq) (*pb.AccessCountRes, error) {
	res, err := s.storage.DashboardS.GetDailyPersonalAccessCount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func (s *DashboardService) GetDailyPersonalBookingRevenueByDay(ctx context.Context, req *pb.BookingRevenueReq) (*pb.BookingRevenueRes, error) {
	res, err := s.storage.DashboardS.GetDailyPersonalBookingRevenueByDay(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *DashboardService) GetDailyAccessCountBySubscriptionID(ctx context.Context, req *pb.SubscriptionCountReq) (*pb.SubscriptionCountRes, error) {
	res, err := s.storage.DashboardS.GetDailyAccessCountBySubscriptionID(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *DashboardService) GetDailyBookingRevenueBySubscriptionID(ctx context.Context, req *pb.BookingRevenueBySubscriptionReq) (*pb.BookingRevenueBySubscriptionRes, error) {
	res, err := s.storage.DashboardS.GetDailyBookingRevenueBySubscriptionID(req)
    if err != nil {
        return nil, err
    }
    return res, nil
}

func (s *DashboardService) TotalMen(ctx context.Context, req *pb.TotalMenReq) (*pb.TotalMenRes, error) {
	res, err := s.storage.DashboardS.TotalMen(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) TotalWomen(ctx context.Context, req *pb.TotalWomenReq) (*pb.TotalWomenRes, error) {
	res, err := s.storage.DashboardS.TotalWomen(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) TotalMembers(ctx context.Context, req *pb.TotalMembersReq) (*pb.TotalMembersRes, error) {
	res, err := s.storage.DashboardS.TotalMembers(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) TotalAmount(ctx context.Context, req *pb.TotalAmountReq) (*pb.TotalAmountRes, error) {
	res, err := s.storage.DashboardS.TotalAmount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
