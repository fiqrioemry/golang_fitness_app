package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:text;not null" json:"-"`
	Role      string         `gorm:"type:varchar(255);default:'customer';check:role IN ('customer','admin','instructor')" json:"role"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Profile      Profile       `gorm:"foreignKey:UserID" json:"profile"`
	Tokens       []Token       `gorm:"foreignKey:UserID" json:"-"`
	Payments     []Payment     `gorm:"foreignKey:UserID" json:"payments,omitempty"`
	UserPackages []UserPackage `gorm:"foreignKey:UserID" json:"packages,omitempty"`
	Bookings     []Booking     `gorm:"foreignKey:UserID" json:"bookings,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

type Token struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36);index;not null" json:"userId"`
	Token     string         `gorm:"type:text;not null" json:"token"`
	ExpiredAt time.Time      `json:"expiredAt"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

type Profile struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36);uniqueIndex;not null" json:"userId"`
	Fullname  string         `gorm:"type:varchar(255);not null" json:"fullname"`
	Birthday  *time.Time     `json:"birthday,omitempty"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	Gender    string         `gorm:"type:varchar(10)" json:"gender"`
	Avatar    string         `gorm:"type:varchar(255)" json:"avatar"`
	Bio       string         `gorm:"type:text" json:"bio"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

type Class struct {
	ID             uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`
	Image          string    `gorm:"type:varchar(255);not null" json:"image"`
	IsActive       bool      `gorm:"default:true" json:"isActive"`
	Duration       int       `gorm:"not null" json:"duration"`
	Description    string    `gorm:"type:text" json:"description"`
	Additional     string    `gorm:"type:longtext" json:"-"`
	AdditionalList []string  `gorm:"-" json:"additional"`
	TypeID         uuid.UUID `json:"typeId"`
	LevelID        uuid.UUID `json:"levelId"`
	LocationID     uuid.UUID `json:"locationId"`
	CategoryID     uuid.UUID `json:"categoryId"`
	SubcategoryID  uuid.UUID `json:"subcategoryId"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"createdAt"`

	// relationship
	Type        Type        `gorm:"foreignKey:TypeID"`
	Level       Level       `gorm:"foreignKey:LevelID"`
	Category    Category    `gorm:"foreignKey:CategoryID"`
	Subcategory Subcategory `gorm:"foreignKey:SubcategoryID"`
	Location    Location    `gorm:"foreignKey:LocationID"`

	// optional
	Galleries []*ClassGallery `gorm:"foreignKey:ClassID" json:"galleries,omitempty"`
	Reviews   []Review        `gorm:"foreignKey:ClassID" json:"reviews,omitempty"`
}

type ClassGallery struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID   uuid.UUID `gorm:"type:char(36);not null;index" json:"classId"`
	URL       string    `gorm:"type:varchar(255);not null" json:"imageUrl"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func (c *Class) BeforeSave(tx *gorm.DB) (err error) {
	if c.AdditionalList != nil {
		jsonBytes, err := json.Marshal(c.AdditionalList)
		if err != nil {
			return err
		}
		c.Additional = string(jsonBytes)
	}
	return nil
}

func (c *Class) AfterFind(tx *gorm.DB) (err error) {
	if c.Additional != "" {
		var tags []string
		if err := json.Unmarshal([]byte(c.Additional), &tags); err != nil {
			return err
		}
		c.AdditionalList = tags
	}
	return nil
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}

type Package struct {
	ID             uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Description    string    `gorm:"type:text;not null" json:"description"`
	IsActive       bool      `gorm:"default:true" json:"isActive"`
	Image          string    `gorm:"type:varchar(255)" json:"image"`
	Price          float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Credit         int       `gorm:"not null" json:"credit"`
	Expired        *int      `json:"expired"`
	Additional     string    `gorm:"type:longtext" json:"-"`
	AdditionalList []string  `gorm:"-" json:"additional"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (p *Package) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}
func (p *Package) BeforeSave(tx *gorm.DB) error {
	if p.AdditionalList != nil {
		data, err := json.Marshal(p.AdditionalList)
		if err != nil {
			return err
		}
		p.Additional = string(data)
	}
	return nil
}

func (p *Package) AfterFind(tx *gorm.DB) error {
	if p.Additional != "" {
		var data []string
		if err := json.Unmarshal([]byte(p.Additional), &data); err != nil {
			return err
		}
		p.AdditionalList = data
	}
	return nil
}

type PackageClass struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID   uuid.UUID `gorm:"type:char(36);not null" json:"classId"`
	PackageID uuid.UUID `gorm:"type:char(36);not null" json:"packageId"`
}

func (pc *PackageClass) BeforeCreate(tx *gorm.DB) (err error) {
	if pc.ID == uuid.Nil {
		pc.ID = uuid.New()
	}
	return
}

type UserPackage struct {
	ID              uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	UserID          uuid.UUID  `gorm:"type:char(36);not null" json:"userId"`
	PackageID       uuid.UUID  `gorm:"type:char(36);not null" json:"packageId"`
	RemainingCredit int        `gorm:"not null;default:0" json:"remainingCredit"`
	ExpiredAt       *time.Time `json:"expiredAt"`
	PurchasedAt     time.Time  `gorm:"autoCreateTime" json:"purchasedAt"`

	User    User    `gorm:"foreignKey:UserID" json:"user"`
	Package Package `gorm:"foreignKey:PackageID" json:"package"`
}

func (up *UserPackage) BeforeCreate(tx *gorm.DB) (err error) {
	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}
	return
}

type Payment struct {
	ID            uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	PackageID     uuid.UUID `gorm:"type:char(36);not null" json:"packageId"`
	UserID        uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	PaymentMethod string    `gorm:"type:varchar(50);not null" json:"paymentMethod"`
	Status        string    `gorm:"type:varchar(20);default:'pending';check:status IN ('success', 'pending', 'failed')" json:"status"`
	PaidAt        time.Time `gorm:"autoCreateTime" json:"paidAt"`

	Package Package `gorm:"foreignKey:PackageID" json:"package"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

type ClassSchedule struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID      uuid.UUID `gorm:"type:char(36);not null" json:"classId"`
	InstructorID uuid.UUID `gorm:"type:char(36);not null" json:"instructorId"`
	Capacity     int       `gorm:"not null" json:"capacity"`
	IsActive     bool      `gorm:"default:true" json:"isActive"`
	StartTime    time.Time `gorm:"not null" json:"startTime"`
	EndTime      time.Time `gorm:"not null" json:"endTime"`

	Class      Class      `gorm:"foreignKey:ClassID" json:"class"`
	Instructor Instructor `gorm:"foreignKey:InstructorID" json:"instructor"`

	BookedCount int `gorm:"-" json:"bookedCount"`
}

func (cs *ClassSchedule) BeforeCreate(tx *gorm.DB) (err error) {
	if cs.ID == uuid.Nil {
		cs.ID = uuid.New()
	}
	return
}

type Booking struct {
	ID              uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID          uuid.UUID      `gorm:"type:char(36);not null" json:"userId"`
	ClassScheduleID uuid.UUID      `gorm:"type:char(36);not null" json:"classScheduleId"`
	Status          string         `gorm:"type:varchar(50);not null;default:'booked'" json:"status"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	ClassSchedule ClassSchedule `gorm:"foreignKey:ClassScheduleID" json:"classSchedule"`
	User          User          `gorm:"foreignKey:UserID" json:"user"`
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}

type Location struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Address     string    `gorm:"type:varchar(255);not null" json:"address"`
	GeoLocation string    `gorm:"type:varchar(255);not null" json:"geoLocation"`
}

func (l *Location) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return
}

type Category struct {
	ID   uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name string    `gorm:"type:varchar(255);not null" json:"name"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}

type Subcategory struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	CategoryID uuid.UUID `gorm:"type:char(36);not null" json:"categoryId"`

	Category Category `gorm:"foreignKey:CategoryID"`
}

func (s *Subcategory) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}

type Type struct {
	ID   uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name string    `gorm:"type:varchar(255);not null" json:"name"`
}

func (t *Type) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

type Level struct {
	ID   uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Name string    `gorm:"type:varchar(255);not null" json:"name"`
}

func (l *Level) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return
}

type Notification struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null;index" json:"userId"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	IsRead    bool           `gorm:"default:false" json:"isRead"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return
}

type Voucher struct {
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Code         string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"code"`
	Description  string         `gorm:"type:text" json:"description"`
	DiscountType string         `gorm:"type:varchar(20);not null" json:"discountType"`
	Discount     float64        `gorm:"not null" json:"discount"`
	MaxDiscount  *float64       `json:"maxDiscount,omitempty"`
	Quota        int            `gorm:"not null" json:"quota"`
	ExpiredAt    time.Time      `gorm:"not null" json:"expiredAt"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type Review struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	ClassID   uuid.UUID `gorm:"type:char(36);not null" json:"classId"`
	Rating    int       `gorm:"not null" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`

	User  User  `gorm:"foreignKey:UserID" json:"user"`
	Class Class `gorm:"foreignKey:ClassID" json:"class"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}

type Attendance struct {
	ID              uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID          uuid.UUID      `gorm:"type:char(36);not null;index" json:"userId"`
	ClassScheduleID uuid.UUID      `gorm:"type:char(36);not null;index" json:"classScheduleId"`
	Status          string         `gorm:"type:varchar(20);not null;check:status IN ('attended', 'absent', 'cancelled')" json:"status"`
	CheckedAt       *time.Time     `json:"checkedAt,omitempty"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	ClassSchedule ClassSchedule `gorm:"foreignKey:ClassScheduleID" json:"classSchedule"`
	User          User          `gorm:"foreignKey:UserID"`
}

type Instructor struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID         uuid.UUID      `gorm:"type:char(36);not null;uniqueIndex" json:"userId"`
	Experience     int            `gorm:"not null;default:0" json:"experience"`
	Specialties    string         `gorm:"type:text" json:"specialties"`
	Rating         float64        `gorm:"type:decimal(2,1);default:5.0" json:"rating"`
	TotalClass     int            `gorm:"not null;default:0" json:"totalClass"`
	Certifications string         `gorm:"type:text" json:"certifications"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UserID"`
}

func (i *Instructor) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return
}

type ScheduleTemplate struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID      uuid.UUID `gorm:"type:char(36);not null" json:"classId"`
	InstructorID uuid.UUID `gorm:"type:char(36);not null" json:"instructorId"`
	DayOfWeek    int       `gorm:"not null" json:"dayOfWeek"`   // 0=Sunday, 6=Saturday
	StartHour    int       `gorm:"not null" json:"startHour"`   // 0-23
	StartMinute  int       `gorm:"not null" json:"startMinute"` // 0-59
	Capacity     int       `gorm:"not null" json:"capacity"`
	IsActive     bool      `gorm:"default:true" json:"isActive"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func (s *ScheduleTemplate) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}
