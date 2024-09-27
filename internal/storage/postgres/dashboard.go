package postgres

import (
	pb "auth-athlevo/genproto/auth"
	"database/sql"
	"fmt"
	"time"

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

var (
	// Time layout for date strings
	dateLayout = "2006-01-02"
)

// GetDailyPersonalAccessCount retrieves the count of access_personal records grouped by day within a date range for a given gym.
func (r *DashboardRepo) GetDailyPersonalAccessCount(req *pb.AccessCountReq) (*pb.AccessCountRes, error) {
	// Generate a list of all dates within the range
	dates := generateDates(parseDateString(req.StartDate, dateLayout), parseDateString(req.EndDate, dateLayout))

	// Construct the query
	query := `
		WITH Subscriptions AS (
			SELECT id
			FROM subscription_personal
			WHERE gym_id = $1
		),
		Bookings AS (
			SELECT id, user_id
			FROM booking_personal
			WHERE subscription_id IN (SELECT id FROM Subscriptions)
		)
		SELECT 
			DATE(ap.date) AS access_date,
			COUNT(DISTINCT bp.user_id) AS user_count
		FROM access_personal ap
		JOIN booking_personal bp ON ap.booking_id = bp.id
		WHERE bp.subscription_id IN (SELECT id FROM Subscriptions) AND ap.date >= $2 AND ap.date <= $3
		GROUP BY access_date
		ORDER BY access_date;
	`

	rows, err := r.db.Query(query, req.GymId, req.StartDate, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("error getting daily personal access count: %w", err)
	}
	defer rows.Close()

	// Create a map to store the results
	accessCounts := make(map[string]int)
	for rows.Next() {
		var (
			accessDate time.Time
			userCount  int
		)

		err := rows.Scan(&accessDate, &userCount)
		if err != nil {
			return nil, fmt.Errorf("error scanning daily access count: %w", err)
		}

		accessCounts[accessDate.Format(time.RFC3339)] = userCount
	}

	// Fill in missing dates with 0 values
	// dailyAccessCounts := make([]*DailyAccessCount, len(dates))
	dailyAccessCounts := &pb.AccessCountRes{}
	for _, date := range dates {
		formattedDate := date.Format(time.RFC3339)
		count, ok := accessCounts[formattedDate]
		if !ok {
			count = 0
		}
		// dailyAccessCounts[i] = &DailyAccessCount{
		// 	AccessDate:  formattedDate,
		// 	AccessCount: count,
		// }

		dailyAccessCounts.AccessCountList = append(dailyAccessCounts.AccessCountList, &pb.AccessCount{
			AccessDate: formattedDate,
			AccessCount: int32(count),
		})
	}

	return dailyAccessCounts, nil
}

// GetDailyPersonalBookingRevenueByDay retrieves the average revenue from personal bookings grouped by day within a date range for a given gym.
func (r *DashboardRepo) GetDailyPersonalBookingRevenueByDay(req *pb.BookingRevenueReq) (*pb.BookingRevenueRes, error) {
	// Generate a list of all dates within the range
	dates := generateDates(parseDateString(req.StartDate, dateLayout), parseDateString(req.EndDate, dateLayout))

	// Construct the query
	query := `
		WITH Subscriptions AS (
			SELECT id
			FROM subscription_personal
			WHERE gym_id = $1
		),
		Bookings AS (
			SELECT id, payment, start_date
			FROM booking_personal
			WHERE subscription_id IN (SELECT id FROM Subscriptions) AND start_date >= $2 AND start_date <= $3
		)
		SELECT 
			DATE(start_date) AS booking_date,
			SUM(payment) AS total_revenue
		FROM Bookings
		GROUP BY booking_date
		ORDER BY booking_date;
	`

	rows, err := r.db.Query(query, req.GymId, req.StartDate, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("error getting average personal booking revenue by day: %w", err)
	}
	defer rows.Close()

	// Create a map to store the results
	bookingRevenues := make(map[string]float64)
	for rows.Next() {
		var (
			bookingDate  time.Time
			totalRevenue float64
		)

		err := rows.Scan(&bookingDate, &totalRevenue)
		if err != nil {
			return nil, fmt.Errorf("error scanning daily revenue: %w", err)
		}

		bookingRevenues[bookingDate.Format(time.RFC3339)] = totalRevenue
	}

	// Fill in missing dates with 0 values
	// dailyRevenues := make([]*DailyRevenue, len(dates))
	dailyRevenues := &pb.BookingRevenueRes{}
    for _, date := range dates {
        formattedDate := date.Format(time.RFC3339)
        revenue, ok := bookingRevenues[formattedDate]
        if!ok {
            revenue = 0
        }
        // dailyRevenues[i] = &DailyRevenue{
        //     BookingDate:    formattedDate,
        //     AverageRevenue: revenue,
        // }

        dailyRevenues.BookingRevenueList = append(dailyRevenues.BookingRevenueList, &pb.BookingRevenue{
            BookingDate:    formattedDate,
            AverageRevenue: float32(revenue),
        })
    }

	return dailyRevenues, nil
}

// GetDailyAccessCountBySubscriptionID retrieves the count of access_personal records for a given subscription ID grouped by day within a date range.
func (r *DashboardRepo) GetDailyAccessCountBySubscriptionID(req *pb.SubscriptionCountReq) (*pb.SubscriptionCountRes, error) {
	// Generate a list of all dates within the range
	dates := generateDates(parseDateString(req.StartDate, dateLayout), parseDateString(req.EndDate, dateLayout))

	// Construct the query
	query := `
		SELECT DATE(ap.date) AS access_date, COUNT(DISTINCT bp.user_id) AS user_count
		FROM access_personal ap
		JOIN booking_personal bp ON ap.booking_id = bp.id
		WHERE bp.subscription_id = $1 AND ap.date >= $2 AND ap.date <= $3
		GROUP BY access_date
		ORDER BY access_date;
	`

	rows, err := r.db.Query(query, req.SubscriptionID, req.StartDate, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("error getting daily access count by subscription ID: %w", err)
	}
	defer rows.Close()

	// Create a map to store the results
	accessCounts := make(map[string]int)
	for rows.Next() {
		var (
			accessDate time.Time
			userCount  int
		)

		err := rows.Scan(&accessDate, &userCount)
		if err != nil {
			return nil, fmt.Errorf("error scanning daily access count: %w", err)
		}

		accessCounts[accessDate.Format(time.RFC3339)] = userCount
	}

	// Fill in missing dates with 0 values
	// dailyAccessCounts := make([]*DailyAccessCount, len(dates))
	dailyAccessCounts := &pb.SubscriptionCountRes{}
    for _, date := range dates {
        formattedDate := date.Format(time.RFC3339)
        count, ok := accessCounts[formattedDate]
        if!ok {
            count = 0
        }
        // dailyAccessCounts[i] = &DailyAccessCount{
        //     AccessDate:  formattedDate,
        //     AccessCount: count,
        // }

        dailyAccessCounts.SubscriptionCountList = append(dailyAccessCounts.SubscriptionCountList, &pb.SubscriptionCount{
            AccessDate: formattedDate,
            AccessCount: int32(count),
        })
    }

	return dailyAccessCounts, nil
}

// GetDailyBookingRevenueBySubscriptionID retrieves the total revenue from personal bookings for a given subscription ID grouped by day within a date range.
func (r *DashboardRepo) GetDailyBookingRevenueBySubscriptionID(req *pb.BookingRevenueBySubscriptionReq) (*pb.BookingRevenueBySubscriptionRes, error) {
	// Generate a list of all dates within the range
	dates := generateDates(parseDateString(req.StartDate, dateLayout), parseDateString(req.EndDate, dateLayout))

	// Construct the query
	query := `
		SELECT DATE(start_date) AS booking_date, SUM(payment) AS total_revenue
		FROM booking_personal
		WHERE subscription_id = $1 AND start_date >= $2 AND start_date <= $3
		GROUP BY booking_date
		ORDER BY booking_date;
	`

	rows, err := r.db.Query(query, req.SubscriptionID, req.StartDate, req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("error getting daily booking revenue by subscription ID: %w", err)
	}
	defer rows.Close()

	// Create a map to store the results
	bookingRevenues := make(map[string]float64)
	for rows.Next() {
		var (
			bookingDate  time.Time
			totalRevenue float64
		)

		err := rows.Scan(&bookingDate, &totalRevenue)
		if err != nil {
			return nil, fmt.Errorf("error scanning daily revenue: %w", err)
		}

		bookingRevenues[bookingDate.Format(time.RFC3339)] = totalRevenue
	}

	// Fill in missing dates with 0 values
	// dailyRevenues := make([]*DailyRevenue, len(dates))
	dailyRevenues := &pb.BookingRevenueBySubscriptionRes{}
    for _, date := range dates {
        formattedDate := date.Format(time.RFC3339)
        revenue, ok := bookingRevenues[formattedDate]
        if!ok {
            revenue = 0
        }
        // dailyRevenues[i] = &DailyRevenue{
        //     BookingDate:    formattedDate,
        //     AverageRevenue: revenue,
        // }

        dailyRevenues.BookingRevenueBySubscriptionList = append(dailyRevenues.BookingRevenueBySubscriptionList, &pb.BookingRevenueBySubscription{
            BookingDate:    formattedDate,
            AverageRevenue: float32(revenue),
        })
    }

	return dailyRevenues, nil
}

// generateDates generates a slice of dates between the start and end dates.
func generateDates(startDate, endDate time.Time) []time.Time {
	var dates []time.Time
	for date := startDate; date.Before(endDate.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
		dates = append(dates, date)
	}
	return dates
}

func parseDateString(dateStr string, layout string) time.Time {
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Printf("error parsing date string: %v\n", err)  // Log the error instead
		return time.Time{}  // Return zero value of time.Time
	}
	return parsedTime
}


func (r *DashboardRepo) TotalMen(req *pb.TotalMenReq) (*pb.TotalMenRes, error) {
    // Query to count male users
    query := `
        SELECT COUNT(*)
        FROM users
        WHERE gender = 'male' AND gym_id = $1;
    `

    var count int32
    err := r.db.QueryRow(query, req.GymId).Scan(&count)
    if err != nil {
        return nil, fmt.Errorf("error getting total men count: %w", err)
    }

    // Create the TotalMen object
    totalMen := &pb.TotalMen{
        TotalMen: count,
        Gender:   "male", // Assuming the gender is fixed for this method
    }

    // Populate the TotalMenRes
    totalMenRes := &pb.TotalMenRes{
        TotalMenList: []*pb.TotalMen{totalMen}, // Assign to the repeated field
    }

    return totalMenRes, nil
}


func (r *DashboardRepo) TotalWomen(req *pb.TotalWomenReq) (*pb.TotalWomenRes, error) {
    query := `
        SELECT COUNT(*)
        FROM users
        WHERE gender = 'female' AND gym_id = $1;
    `

    var count int32
    err := r.db.QueryRow(query, req.GymId).Scan(&count)
    if err != nil {
        return nil, fmt.Errorf("error getting total women count: %w", err)
    }

    // Correct usage based on original Protobuf
    totalWomen := &pb.TotalWomen{
        TotalWomen: count,
        Gender:     "female",  // Assuming you're keeping track of gender
    }

    return &pb.TotalWomenRes{
        TotalWomenList: []*pb.TotalWomen{totalWomen},  // Field name should be 'totalWomenList'
    }, nil
}


func (r *DashboardRepo) TotalMembers(req *pb.TotalMembersReq) (*pb.TotalMembersRes, error) {
    query := `
        SELECT COUNT(*)
        FROM users
        WHERE gym_id = $1;
    `

    var count int32
    err := r.db.QueryRow(query, req.GymId).Scan(&count)
    if err != nil {
        return nil, fmt.Errorf("error getting total members count: %w", err)
    }

    totalMembers := &pb.TotalMembers{
        TotalMembers: count,
    }

    return &pb.TotalMembersRes{
        TotalMembersList: []*pb.TotalMembers{totalMembers},  // Correct field name
    }, nil
}

func (r *DashboardRepo) TotalAmount(req *pb.TotalAmountReq) (*pb.TotalAmountRes, error) {
    query := `
        SELECT SUM(payment)
        FROM (
            SELECT payment FROM booking_personal WHERE gym_id = $1
            UNION ALL
            SELECT payment FROM booking_group WHERE gym_id = $1
            UNION ALL
            SELECT payment FROM booking_coach WHERE gym_id = $1
        ) AS all_bookings;
    `

    var totalAmount float32
    err := r.db.QueryRow(query, req.GymId).Scan(&totalAmount)
    if err != nil {
        return nil, fmt.Errorf("error getting total amount: %w", err)
    }

    totalAmountRecord := &pb.TotalAmount{
        TotalAmount: totalAmount,
    }

    return &pb.TotalAmountRes{
        TotalAmountList: []*pb.TotalAmount{totalAmountRecord}, 
    }, nil
}


func (r *DashboardRepo) CompareCurrentAndPreviousMonthRevenue(req *pb.Void) (*pb.RevenueReq, error) {
	
	query := `
		WITH current_month_revenue AS (
			SELECT 
				SUM(payment) AS total_revenue
			FROM (
				SELECT payment, created_at FROM booking_personal
				UNION ALL
				SELECT payment, created_at FROM booking_group
				UNION ALL
				SELECT payment, created_at FROM booking_coach
			) AS all_bookings
			WHERE EXTRACT(MONTH FROM created_at) = EXTRACT(MONTH FROM CURRENT_DATE)
			AND EXTRACT(YEAR FROM created_at) = EXTRACT(YEAR FROM CURRENT_DATE)
		),
		previous_month_revenue AS (
			SELECT 
				SUM(payment) AS total_revenue
			FROM (
				SELECT payment, created_at FROM booking_personal
				UNION ALL
				SELECT payment, created_at FROM booking_group
				UNION ALL
				SELECT payment, created_at FROM booking_coach
			) AS all_bookings
			WHERE EXTRACT(MONTH FROM created_at) = EXTRACT(MONTH FROM CURRENT_DATE) - 1
			AND EXTRACT(YEAR FROM created_at) = EXTRACT(YEAR FROM CURRENT_DATE)
		)
		SELECT 
			CASE 
				WHEN previous_month_revenue.total_revenue = 0 THEN 0
				ELSE (current_month_revenue.total_revenue - previous_month_revenue.total_revenue) / previous_month_revenue.total_revenue * 100
			END AS revenue_change_percentage
		FROM current_month_revenue, previous_month_revenue;
	`

	percentageChange := pb.RevenueReq{}

	err := r.db.QueryRow(query).Scan(&percentageChange)
	if err != nil {
		return nil, fmt.Errorf("failed to get revenue change percentage: %w", err)
	}

	return &percentageChange, nil
}


func (r *DashboardRepo) GetMonthlyRevenueStats(req *pb.Void) (*pb.MonthlyRevenueRes, error) {
	query := `
		SELECT 
			EXTRACT(YEAR FROM created_at) AS year,
			EXTRACT(MONTH FROM created_at) AS month,
			SUM(payment) AS total_revenue
		FROM (
			SELECT payment, created_at FROM booking_personal
			UNION ALL
			SELECT payment, created_at FROM booking_group
			UNION ALL
			SELECT payment, created_at FROM booking_coach
		) AS all_bookings
		GROUP BY EXTRACT(YEAR FROM created_at), EXTRACT(MONTH FROM created_at)
		ORDER BY year, month;
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly revenue statistics: %w", err)
	}
	defer rows.Close()

	revenues := &pb.MonthlyRevenueRes{}
	for rows.Next() {
		revenue := pb.MonthlyRevenue{}
		err := rows.Scan(&revenue.Year, &revenue.Month, &revenue.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan monthly revenue: %w", err)
		}
		revenues.MonthlyRevenue = append(revenues.MonthlyRevenue, &revenue)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return revenues, nil
}