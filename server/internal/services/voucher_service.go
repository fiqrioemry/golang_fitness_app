package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type VoucherService interface {
	DecreaseQuota(userID uuid.UUID, code string) error
	CreateVoucher(dto.CreateVoucherRequest) error
	GetAllVouchers() ([]dto.VoucherResponse, error)
	ApplyVoucher(req dto.ApplyVoucherRequest) (*dto.ApplyVoucherResponse, error)
}

type voucherService struct {
	repo repositories.VoucherRepository
}

func NewVoucherService(repo repositories.VoucherRepository) VoucherService {
	return &voucherService{
		repo: repo,
	}
}

func (s *voucherService) CreateVoucher(req dto.CreateVoucherRequest) error {
	expiredAt, err := time.Parse("2006-01-02", req.ExpiredAt)
	if err != nil {
		return err
	}

	voucher := models.Voucher{
		ID:           uuid.New(),
		Code:         req.Code,
		Description:  req.Description,
		DiscountType: req.DiscountType,
		Discount:     req.Discount,
		MaxDiscount:  req.MaxDiscount,
		Quota:        req.Quota,
		ExpiredAt:    expiredAt,
		CreatedAt:    time.Now(),
	}

	err = s.repo.Create(&voucher)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "code") {
			return errors.New("voucher code must be unique")
		}
		return err
	}

	return nil
}

func (s *voucherService) GetAllVouchers() ([]dto.VoucherResponse, error) {
	vouchers, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []dto.VoucherResponse
	for _, v := range vouchers {
		result = append(result, dto.VoucherResponse{
			ID:           v.ID.String(),
			Code:         v.Code,
			Description:  v.Description,
			DiscountType: v.DiscountType,
			Discount:     v.Discount,
			MaxDiscount:  v.MaxDiscount,
			Quota:        v.Quota,
			ExpiredAt:    v.ExpiredAt.Format("2006-01-02"),
			CreatedAt:    v.CreatedAt.Format("2006-01-02"),
		})
	}

	return result, nil
}

func (s *voucherService) ApplyVoucher(req dto.ApplyVoucherRequest) (*dto.ApplyVoucherResponse, error) {
	voucher, err := s.repo.GetValidVoucherByCode(req.Code)
	if err != nil {
		return nil, errors.New("invalid or expired voucher")
	}

	var userUUID uuid.UUID
	hasUser := req.UserID != nil && *req.UserID != ""

	if hasUser {
		userUUID, err = uuid.Parse(*req.UserID)
		if err != nil {
			return nil, errors.New("invalid user id")
		}

		if !voucher.IsReusable {
			used, err := s.repo.CheckVoucherUsed(userUUID, voucher.ID)
			if err != nil {
				return nil, err
			}
			if used {
				return nil, errors.New("voucher already used")
			}
		}
	}

	var discountValue float64
	if voucher.DiscountType == "percentage" {
		discountValue = req.Total * (voucher.Discount / 100)
		if voucher.MaxDiscount != nil && discountValue > *voucher.MaxDiscount {
			discountValue = *voucher.MaxDiscount
		}
	} else {
		discountValue = min(voucher.Discount, req.Total)
	}

	final := req.Total - discountValue

	return &dto.ApplyVoucherResponse{
		Code:          voucher.Code,
		DiscountType:  voucher.DiscountType,
		Discount:      voucher.Discount,
		MaxDiscount:   voucher.MaxDiscount,
		DiscountValue: discountValue,
		FinalTotal:    final,
	}, nil
}

func (s *voucherService) DecreaseQuota(userID uuid.UUID, code string) error {
	voucher, err := s.repo.GetByCode(code)
	if err != nil {
		return err
	}

	if !voucher.IsReusable {
		_ = s.repo.InsertUsedVoucher(userID, voucher.ID)
	}

	if voucher.Quota > 0 {
		voucher.Quota -= 1
		return s.repo.UpdateVoucher(voucher)
	}
	return nil
}
