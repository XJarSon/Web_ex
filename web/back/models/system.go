package models

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type System struct {
	CreateTime   time.Time `json:"createTime";gorm:"column:create_time;primary_key"`
	Version      string    `json:"version";gorm:"column:version"`
	VisitsNumber uint      `json:"visitsNumber";gorm:"column:visits_number"`
}

type Population struct {
	FacultyNumber uint `json:"facultyNumber";gorm:"column:faculty_number"`
	CourseNumber  uint `json:"courseNumber";gorm:"column:course_number"`
}

func GetSystemMessage() (System, error) {
	var system System
	err := db.First(&system).Error
	return system, err
}

func UpdateVisitsNumber() error {
	var system System
	mutex.Lock()
	defer mutex.Unlock()
	err := db.Debug().First(&system).Error
	fmt.Println(system)
	if err != nil {
		return err
	}
	system.VisitsNumber++
	fmt.Println(system)
	err = db.Debug().Model(&system).Where("create_time = ?", system.CreateTime).Update("visits_number", system.VisitsNumber).Error
	return err
}

func GetPopulationMessage() Population {
	var population Population
	db.Raw("Select count(*) as faculty_number from faculties").Scan(&population.FacultyNumber)
	db.Raw("Select count(*) as course_number from courses").Scan(&population.CourseNumber)
	return population
}
