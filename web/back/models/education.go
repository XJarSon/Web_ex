package models

// Faculty 对学院、专业、班级、学生的增删改查
type Faculty struct {
	// 学院ID
	FID uint `json:"fid" gorm:"primaryKey;autoIncrement"`
	// 学院代码
	FCODE string `json:"fcode" gorm:"unique;not null;type:char(2)"`
	// 学院名称
	Name string `json:"name" gorm:"not null;varchar(10)"`
}

type Course struct {
	// 课程ID
	CID uint `json:"cid" gorm:"primaryKey;autoIncrement"`
	// 课程名
	Title string `json:"title" gorm:"not null;varchar(10)"`
	// 课时
	Hours uint `json:"hours" gorm:"not null"`
	// 封面
	Cover string `json:"cover" gorm:"not null"`
}

type FC struct {
	FID uint `json:"fid" gorm:"column:f_id"`
	CID uint `json:"cid" gorm:"column:c_id"`
}

// GetCourse  获取课程信息
func GetCourse(title string) Course {
	var course Course
	if err := db.First(&course, "title = ?", title).Error; err != nil {
		return Course{}
	} else {
		return course
	}
}

func GetAllCourses() []Course {
	var courses []Course
	db.Find(&courses)
	return courses
}
func GetFacultyByID(FID uint) Faculty {
	var faculty Faculty

	if err := db.First(&faculty, FID).Error; err != nil {
		return Faculty{}
	} else {
		return faculty
	}
}

func AddFaculty(faculty Faculty) uint {
	if err := db.Omit("f_id").Create(&faculty).Error; err != nil {
		return 0
	} else {
		return faculty.FID
	}
}

func DeleteFaculty(FID uint) error {
	if err := db.Delete(&Faculty{}, FID).Error; err != nil {
		return err
	} else {
		return nil
	}

}

func AddCourse(course *Course) error {
	if err := db.Omit("c_id").Create(&course).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteCourse(CID uint) error {
	if err := db.Delete(&Course{}, CID).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func AddFC(fc *FC) error {
	if err := db.Table("fcs").Create(&fc).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteFC(fc *FC) error {
	if err := db.Table("fcs").Where("f_id = ? and c_id = ?", fc.FID, fc.CID).Delete(&fc).Error; err != nil {
		return err
	} else {
		return nil
	}
}
