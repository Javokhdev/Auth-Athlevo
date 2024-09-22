package postgres

import (
	pb "auth-athlevo/genproto/auth"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DashboardRepo implements the DashboardRepoI interface for dashboard statistics.
type DashboardRepo struct {
	db *sql.DB
}

// NewDashboardRepo creates a new DashboardRepo.
func NewDashboardRepo(db *sql.DB) *DashboardRepo {
	return &DashboardRepo{
		db: db,
	}
}

// GetPersonalAccessCount retrieves the count of access_personal records within a date range for a given gym.
func (r *DashboardRepo) GetPersonalAccessCount(req *pb.AccessCountReq) (*pb.AccessCountRes, error) {
	res := &pb.AccessCountRes{}

	query := `
		WITH Subscriptions AS (
			SELECT id
			FROM subscription_personal
			WHERE gym_id = $1
		),
		Bookings AS (
			SELECT id
			FROM booking_personal
			WHERE subscription_id IN (SELECT id FROM Subscriptions)
		)
		SELECT COUNT(*)
		FROM access_personal
		WHERE booking_id IN (SELECT id FROM Bookings) AND date >= $2 AND date <= $3;
	`

	err := r.db.QueryRow(query, req.GymId, req.StartDate, req.EndDate).Scan(&res.Count)
	if err != nil {
		return nil, fmt.Errorf("error getting personal access count: %w", err)
	}

	return res, nil
}

// GetTotalPersonalBookingRevenue retrieves the total revenue from personal bookings within a date range for a given gym.
func (r *DashboardRepo) GetTotalPersonalBookingRevenue(req *pb.TotalRevenueReq) (*pb.TotalRevenueRes, error) {
	res := &pb.TotalRevenueRes{}

	query := `
		WITH Subscriptions AS (
			SELECT id
			FROM subscription_personal
			WHERE gym_id = $1
		),
		Bookings AS (
			SELECT id, payment
			FROM booking_personal
			WHERE subscription_id IN (SELECT id FROM Subscriptions) AND start_date >= $2 AND start_date <= $3
		)
		SELECT SUM(payment)
		FROM Bookings;
	`

	err := r.db.QueryRow(query, req.GetGymId, req.StartDate, req.EndDate).Scan(&res.TotalRevenue)
	if err != nil {
		return nil, fmt.Errorf("error getting total personal booking revenue: %w", err)
	}

	return res, nil
}

// GetAccessCountBySubscriptionID retrieves the count of access_personal records for a given subscription ID within a date range.
func (r *DashboardRepo) GetAccessCountBySubscriptionID(req *pb.SubscriptionCountReq) (*pb.SubscriptionCountRes, error) {
	res := &pb.SubscriptionCountRes{}

	query := `
		SELECT COUNT(*)
		FROM access_personal ap
		JOIN booking_personal bp ON ap.booking_id = bp.id
		WHERE bp.subscription_id = $1 AND ap.date >= $2 AND ap.date <= $3;
	`

	err := r.db.QueryRow(query, req.SubscriptionID, req.StartDate, req.EndDate).Scan(&res.Count)
	if err != nil {
		return nil, fmt.Errorf("error getting access count by subscription ID: %w", err)
	}

	return res, nil
}

// GetBookingRevenueBySubscriptionID retrieves the total revenue from personal bookings for a given subscription ID within a date range.
func (r *DashboardRepo) GetBookingRevenueBySubscriptionID(req *pb.TotalRevenueBySubscriptionReq) (*pb.TotalRevenueBySubscriptionRes, error) {
	res := &pb.TotalRevenueBySubscriptionRes{}

	query := `
		SELECT SUM(payment)
		FROM booking_personal
		WHERE subscription_id = $1 AND start_date >= $2 AND start_date <= $3;
	`

	err := r.db.QueryRow(query, req.SubscriptionID, req.StartDate, req.EndDate).Scan(&res.TotalRevenue)
	if err != nil {
		return nil, fmt.Errorf("error getting booking revenue by subscription ID: %w", err)
	}

	return res, nil
}
