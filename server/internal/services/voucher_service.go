package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type VoucherService interface {
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

	return s.repo.Create(&voucher)
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

	var discountValue float64
	if voucher.DiscountType == "percentage" {
		discountValue = req.Total * (voucher.Discount / 100)
		if voucher.MaxDiscount != nil && discountValue > *voucher.MaxDiscount {
			discountValue = *voucher.MaxDiscount
		}
	} else if voucher.DiscountType == "fixed" {
		discountValue = voucher.Discount
		if discountValue > req.Total {
			discountValue = req.Total
		}
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
