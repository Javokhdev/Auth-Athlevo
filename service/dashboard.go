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

// func (s *DashboardService) GetDailyPersonalAccessCount(ctx context.Context, req *pb.AccessCountReq) (*pb.AccessCountRes, error) {
// 	res, err := s.storage.DashboardS.GetDailyPersonalAccessCount(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }


// func (s *DashboardService) GetDailyPersonalBookingRevenueByDay(ctx context.Context, req *pb.BookingRevenueReq) (*pb.BookingRevenueRes, error) {
// 	res, err := s.storage.DashboardS.GetDailyPersonalBookingRevenueByDay(req)
//     if err != nil {
//         return nil, err
//     }
//     return res, nil
// }

// func (s *DashboardService) GetDailyAccessCountBySubscriptionID(ctx context.Context, req *pb.SubscriptionCountReq) (*pb.SubscriptionCountRes, error) {
// 	res, err := s.storage.DashboardS.GetDailyAccessCountBySubscriptionID(req)
//     if err != nil {
//         return nil, err
//     }
//     return res, nil
// }

// func (s *DashboardService) GetDailyBookingRevenueBySubscriptionID(ctx context.Context, req *pb.BookingRevenueBySubscriptionReq) (*pb.BookingRevenueBySubscriptionRes, error) {
// 	res, err := s.storage.DashboardS.GetDailyBookingRevenueBySubscriptionID(req)
//     if err != nil {
//         return nil, err
//     }
//     return res, nil
// }

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

func (s *DashboardService) CompareCurrentAndPreviousMonthRevenue(ctx context.Context, req *pb.Void) (*pb.RevenueReq, error) {
	res, err := s.storage.DashboardS.CompareCurrentAndPreviousMonthRevenue(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) GetMonthlyRevenueStats(ctx context.Context, req *pb.Void) (*pb.MonthlyRevenueRes, error) {
	res, err := s.storage.DashboardS.GetMonthlyRevenueStats(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) GetGenderCounts(ctx context.Context, req *pb.TotalGenderReq) (*pb.GenderCountsRes, error) {
	res, err := s.storage.DashboardS.GetGenderCounts(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *DashboardService) GetRevenueByTariff(ctx context.Context, req *pb.TotalRevenueReq) (*pb.TariffRevenueRes, error) {
	res, err := s.storage.DashboardS.GetRevenueByTariff(req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}

func (s *DashboardService) GetUsersByTariff(ctx context.Context, req *pb.TotalUsersReq) (*pb.TariffUsersRes, error) {
	res, err := s.storage.DashboardS.GetUsersByTariff(req)
    if err!= nil {
        return nil, err
    }
    return res, nil
}