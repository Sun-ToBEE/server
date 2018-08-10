package models

import (
	"time"
	"github.com/sunyatsuntobee/server/logger"
)

// Organization Model
type Organization struct {
	ID              int       `xorm:"id INT PK NOTNULL UNIQUE AUTOINCR" json:"id"`
	Name            string    `xorm:"name VARCHAR(20) NOTNULL" json:"name"`
	Phone           string    `xorm:"phone VARCHAR(20) NOTNULL UNIQUE" json:"phone"`
	Password        string    `xorm:"password VARCHAR(50) NOTNULL" json:"password"`
	College         int       `xorm:"college INT NOTNULL" json:"college"`
	CollegeDistrict int       `xorm:"college_district INT NOTNULL" json:"college_district"`
	LogoURL         string    `xorm:"logo_url VARCHAR(50)" json:"logo_url"`
	Description     string    `xorm:"description VARCHAR(200)" json:"description"`

	InvitationCode  string    `xorm:"invitation_code VARCHAR(50) NOTNULL" json:"invitation_code"` //邀请码
	SpreadFlag      bool      `xorm:"spread_flag BOOL NOTNULL" json:"spread_flag"` //是否被推广
	SpreadDay       int       `xorm:"spread_day INT" json:"spread_day"`   //推广天数
	SpreadStartTime time.Time `xorm:"spread_start_time DATETIME" json:"spread_start_time"` //推广开始时间
	TobeeCoin       int       `xorm:"tobee_coin INT" json:"tobee_coin"`   //图币
}

// NewOrganization creates a new organization
func NewOrganization(name, phone, password,
	logoURL, description, invitationCode string, spreadFlag bool,
	college, collegeDistrict, spreadDay int, 
	spreadStartTime time.Time) *Organization {
	return &Organization{
		Name:            name,
		Phone:           phone,
		Password:        password,
		College:         college,
		CollegeDistrict: collegeDistrict,
		LogoURL:         logoURL,
		Description:     description,
		InvitationCode:  invitationCode,
		SpreadFlag:      spreadFlag,
		SpreadDay:       spreadDay,
		SpreadStartTime: spreadStartTime,
	}
}

// OrganizationDataAccessObject provides database access for Model
// Organization
type OrganizationDataAccessObject struct{}

// OrganizationDAO instance of OrganizationDataAccessObject
var OrganizationDAO *OrganizationDataAccessObject

// TableName returns table name
func (*OrganizationDataAccessObject) TableName() string {
	return "organizations"
}

// FindAll finds all organizations
func (*OrganizationDataAccessObject) FindAll() []Organization {
	organizations := make([]Organization, 0)
	err := orm.Table(OrganizationDAO.TableName()).Find(&organizations)
	logger.LogIfError(err)
	return organizations
}

// FindFullByID finds all joined organizations according to ID
func (*OrganizationDataAccessObject) FindFullByID(id int) []OrganizationFull {
	organizations := make([]OrganizationFull, 0)
	err := orm.Table(OrganizationDAO.TableName()).
		Join("LEFT OUTER", "organization_contact_relationships",
			"organizations.id=organization_contact_relationships.organization_id").
		Join("LEFT OUTER", UserDAO.TableName(),
			"organization_contact_relationships.contact_id=users.id").
		Join("LEFT OUTER", OrganizationDepartmentDAO.TableName(),
			"organizations.id=organization_departments.organization_id").
		Join("LEFT OUTER", ActivityDAO.TableName(),
			"organizations.id=activities.organization_id").
		Join("LEFT OUTER", ActivityStageDAO.TableName(),
			"activities.id=activity_stages.activity_id").
		Where("organizations.id=?", id).Find(&organizations)
	logger.LogIfError(err)
	return organizations

}

// FindByID finds an organization according to ID
func (*OrganizationDataAccessObject) FindByID(id int) (Organization, bool) {
	var o Organization
	has, err := orm.Table(OrganizationDAO.TableName()).ID(id).Get(&o)
	logger.LogIfError(err)
	return o, has
}

// FindByPhone finds an organization according to Phone
func (*OrganizationDataAccessObject) FindByPhone(phone string) (Organization, bool) {
	var o Organization
	has, err := orm.Table(OrganizationDAO.TableName()).Where("phone=?", phone).Get(&o)
	logger.LogIfError(err)
	return o, has
}

// InsertOne inserts an organization
func (*OrganizationDataAccessObject) InsertOne(organization *Organization) {
	_, err := orm.Table(OrganizationDAO.TableName()).InsertOne(organization)
	logger.LogIfError(err)
}

// UpdateOne updates an organization
func (*OrganizationDataAccessObject) UpdateOne(o *Organization) {
	_, err := orm.Table(OrganizationDAO.TableName()).ID(o.ID).Update(o)
	logger.LogIfError(err)
}

func (*OrganizationDataAccessObject) FindByOpenid(
	openid string) (Organization, bool) {
	var o Organization
	has, err := orm.Table(OrganizationDAO.TableName()).
		Where("openid=?", openid).Get(&o)
	logger.LogIfError(err)
	return o, has
}

// FindByCDID find all organizations by collegedistrictid
func (*OrganizationDataAccessObject) FindByCDID(cdid int) []Organization {
    organizations := make([]Organization, 0)
	err := orm.Table(OrganizationDAO.TableName()).Where("college_district=?", cdid).
	Find(& organizations)
	logger.LogIfError(err)
	return organizations;
}