package storage

import pb "auth-athlevo/genproto/auth"

type StorageI interface {
	Auth() AuthI
	User() UserI
    Dashboard() DashboardI
}

type AuthI interface {
	Register(*pb.RegisterReq) (*pb.RegisterRes, error)
    Login(*pb.LoginReq) (*pb.User, error)
    ForgotPassword(*pb.GetByEmail) (*pb.ForgotPassRes, error)
    ResetPassword(*pb.ResetPassReq) (*pb.ResetPasswordRes, error)
    SaveRefreshToken(*pb.RefToken) (*pb.SaveRefereshTokenRes, error)
    RefreshToken(*pb.GetByEmail) (*pb.LoginRes, error)
    ChangeRole(*pb.Role) (*pb.ChangeRoleRes, error)
}

type UserI interface {
	GetProfile(*pb.GetByIdReq) (*pb.UserRepeated, error)
    EditProfile(*pb.UserRes) (*pb.UserRes, error)
    ChangePassword(*pb.ChangePasswordReq) (*pb.ChangePasswordRes, error)
    GetSetting(*pb.GetById) (*pb.Setting, error)
    EditSetting(*pb.SettingReq) (*pb.SettingRes, error)
    DeleteUser(*pb.GetById) (*pb.DeleteRes, error)
}

type DashboardI interface { 
    // GetDailyPersonalAccessCount(*pb.AccessCountReq) (*pb.AccessCountRes, error)
    // GetDailyPersonalBookingRevenueByDay(*pb.BookingRevenueReq) (*pb.BookingRevenueRes, error)
    // GetDailyAccessCountBySubscriptionID(*pb.SubscriptionCountReq) (*pb.SubscriptionCountRes, error)
    // GetDailyBookingRevenueBySubscriptionID(*pb.BookingRevenueBySubscriptionReq) (*pb.BookingRevenueBySubscriptionRes, error)
    TotalMen(*pb.TotalMenReq) (*pb.TotalMenRes, error)
    TotalWomen(*pb.TotalWomenReq) (*pb.TotalWomenRes, error)
    TotalMembers(*pb.TotalMembersReq) (*pb.TotalMembersRes, error)
    TotalAmount(*pb.TotalAmountReq) (*pb.TotalAmountRes, error)
    CompareCurrentAndPreviousMonthRevenue(*pb.Void) (*pb.RevenueReq, error)
    GetMonthlyRevenueStats(*pb.Void) (*pb.MonthlyRevenueRes, error)
    GetGenderCounts(*pb.TotalGenderReq) (*pb.GenderCountsRes, error)
    GetRevenueByTariff(*pb.Void) (*pb.TariffRevenueRes, error)
    GetUsersByTariff(*pb.Void) (*pb.TariffUsersRes, error)
}