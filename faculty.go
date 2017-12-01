package main

import (
	"net/http"
	"fmt"
	"Starfleet/global"
)

func facultyViewSchedule(w http.ResponseWriter,r *http.Request){
	_, user := CheckLoginStatus(w,r)
	type FacultySchedule struct{
		CourseName string
		BuildingName string
		RoomNumber string
		MeetingDay string
		Time string
		Year int
		Season string
	}

	facultySchedule := []FacultySchedule{}
	db.Raw(`
	SELECT course_name, building_name,room_number,meeting_day,time,year,season
	FROM section
	JOIN course ON course.course_id = section.course_id
	JOIN time_slot ON time_slot.time_slot_id = section.time_slot_id
	JOIN semester ON time_slot.semester_id = semester.semester_id
	JOIN period ON time_slot.period_id = period.period_id
	JOIN day ON time_slot.day_id = day.day_id
	JOIN location ON section.location_id = location.location_id
	JOIN building ON location.building_id = building.building_id
	JOIN room ON location.room_id = room.room_id
	WHERE section.faculty_id = ? AND semester.year = 2018 AND semester.season = 'Spring'`, user.UserID).Scan(&facultySchedule)
	fmt.Println(facultySchedule)

	err := global.Tpl.ExecuteTemplate(w, "ViewFacultyScheduleDetails", facultySchedule)
	if err != nil {
		fmt.Println(err.Error())
	}
}
