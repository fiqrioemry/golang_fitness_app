package services

import (
	"mime/multipart"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"
)

type ProfileService interface {
	GetUserByID(userID string) (*models.User, error)
	UpdateProfile(userID string, req dto.UpdateProfileRequest) error
	UpdateAvatar(userID string, file *multipart.FileHeader) (string, error)
	GetUserTransactions(userID string, page, limit int) (*dto.TransactionListResponse, error)
	GetUserPackages(userID string, page, limit int) (*dto.UserPackageListResponse, error)
	GetUserPackagesByClassID(userID, classID string) ([]dto.UserPackageResponse, error)
	GetUserBookings(userID string, page, limit int) (*dto.BookingListResponse, error)
}

type profileService struct {
	repo repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) ProfileService {
	return &profileService{repo}
}

func (s *profileService) GetUserByID(userID string) (*models.User, error) {
	return s.repo.GetUserByID(userID)
}
func (s *profileService) UpdateProfile(userID string, req dto.UpdateProfileRequest) error {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	user.Profile.Fullname = req.Fullname
	user.Profile.Gender = req.Gender
	user.Profile.Phone = req.Phone
	user.Profile.Bio = req.Bio
	if req.Birthday != "" {
		birthday, err := time.Parse("2006-01-02", req.Birthday)
		if err == nil {
			user.Profile.Birthday = &birthday
		}
	}

	return s.repo.UpdateUser(user)
}

func (s *profileService) UpdateAvatar(userID string, file *multipart.FileHeader) (string, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return "", err
	}

	if file == nil {
		return "", nil
	}

	if err := utils.ValidateImageFile(file); err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	newAvatarURL, err := utils.UploadToCloudinary(src)
	if err != nil {
		return "", err
	}

	if user.Profile.Avatar != "" && user.Profile.Avatar != newAvatarURL && !isDiceBear(user.Profile.Avatar) {
		_ = utils.DeleteFromCloudinary(user.Profile.Avatar)
	}

	user.Profile.Avatar = newAvatarURL
	err = s.repo.UpdateUser(user)
	return newAvatarURL, err
}

func isDiceBear(url string) bool {
	return url != "" && (len(url) > 0 && (url[:30] == "https://api.dicebear.com" || url[:31] == "https://avatars.dicebear.com"))
}

func (s *profileService) GetUserTransactions(userID string, page, limit int) (*dto.TransactionListResponse, error) {
	offset := (page - 1) * limit

	payments, total, err := s.repo.GetUserTransactions(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(payments) == 0 {
		payments = make([]models.Payment, 0)
	}

	taxRate := utils.GetTaxRate()
	var transactions []dto.TransactionResponse
	for _, p := range payments {
		transactions = append(transactions, dto.TransactionResponse{
			ID:            p.ID.String(),
			PackageID:     p.PackageID.String(),
			PackageName:   p.Package.Name,
			PaymentMethod: p.PaymentMethod,
			Status:        p.Status,
			BasePrice:     p.BasePrice,
			Tax:           taxRate,
			Price:         p.Total,
			PaidAt:        p.PaidAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &dto.TransactionListResponse{
		Transactions: transactions,
		Total:        total,
		Page:         page,
		Limit:        limit,
	}, nil
}

func (s *profileService) GetUserPackages(userID string, page, limit int) (*dto.UserPackageListResponse, error) {
	offset := (page - 1) * limit
	pkgs, total, err := s.repo.GetUserPackages(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(pkgs) == 0 {
		pkgs = make([]models.UserPackage, 0)
	}

	var responses []dto.UserPackageResponse
	for _, p := range pkgs {
		expired := ""
		if p.ExpiredAt != nil {
			expired = p.ExpiredAt.Format("2006-01-02")
		}
		responses = append(responses, dto.UserPackageResponse{
			ID:              p.ID.String(),
			PackageName:     p.Package.Name,
			RemainingCredit: p.RemainingCredit,
			ExpiredAt:       expired,
			PurchasedAt:     p.PurchasedAt.Format("2006-01-02"),
		})
	}

	return &dto.UserPackageListResponse{
		Packages: responses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

func (s *profileService) GetUserPackagesByClassID(userID, classID string) ([]dto.UserPackageResponse, error) {
	userPackages, err := s.repo.GetUserPackagesByClassID(userID, classID)
	if err != nil {
		return nil, err
	}

	var result []dto.UserPackageResponse
	for _, up := range userPackages {
		var (
			expiredAt     string
			expiredInDays int
		)

		if up.ExpiredAt != nil {
			expiredAt = up.ExpiredAt.Format("2006-01-02")
			expiredInDays = int(max(0, int(time.Until(*up.ExpiredAt).Hours()/24)))
		}

		result = append(result, dto.UserPackageResponse{
			ID:              up.ID.String(),
			PackageID:       up.Package.ID.String(),
			PackageName:     up.Package.Name,
			RemainingCredit: up.RemainingCredit,
			ExpiredAt:       expiredAt,
			ExpiredInDays:   expiredInDays,
			PurchasedAt:     up.PurchasedAt.Format("2006-01-02"),
		})
	}

	return result, nil
}

func (s *profileService) GetUserBookings(userID string, page, limit int) (*dto.BookingListResponse, error) {
	offset := (page - 1) * limit
	bookings, total, err := s.repo.GetUserBookings(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(bookings) == 0 {
		bookings = make([]models.Booking, 0)
	}

	var responses []dto.BookingResponse
	for _, b := range bookings {
		cs := b.ClassSchedule
		c := cs.Class
		loc := cs.Class.Location
		instructor := cs.Instructor.User.Profile

		responses = append(responses, dto.BookingResponse{
			ID:             b.ID.String(),
			Status:         b.Status,
			BookedAt:       b.CreatedAt.Format("2006-01-02 15:04:05"),
			ClassID:        c.ID.String(),
			ClassName:      c.Title,
			ClassImage:     c.Image,
			Duration:       c.Duration,
			Date:           cs.Date.Format("2006-01-02"),
			StartHour:      cs.StartHour,
			StartMinute:    cs.StartMinute,
			Location:       loc.Name,
			InstructorName: instructor.Fullname,
			Participant:    cs.Booked,
		})
	}

	return &dto.BookingListResponse{
		Bookings: responses,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}
