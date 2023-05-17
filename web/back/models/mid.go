package models

type ViewFaculty struct {
	// 学院ID
	FID uint `json:"fid" gorm:"primaryKey;autoIncrement"`
	// 学院名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

func GetAllFaculties() ([]*ViewFaculty, error) {
	faculties := make([]*ViewFaculty, 0)
	err := db.Model(&Faculty{}).Find(&faculties).Error
	return faculties, err
}
func GetCoursesByFaculties(fid uint) ([]Course, error) {
	courses := make([]Course, 0)
	err := db.Raw("select courses.c_id c_id ,courses.title title from courses,fcs where courses.c_id = fcs.c_id and fcs.f_id = ?", fid).Scan(&courses).Error
	return courses, err
}
