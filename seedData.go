package main
/*
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"Starfleet/model"
	"os"
	"fmt"
	"time"
)

func main(){
	dbPassword := os.Getenv("PG_DATABASE_PW")
	db, err := gorm.Open("postgres", "host=127.0.0.1 dbname=Starfleet sslmode=disable password="+dbPassword)
	if err != nil {
		fmt.Println("Cannot connect to database...")
		fmt.Println("DB Error: ", err)
	}
	db.SingularTable(true)

	db.DropTable(&model.Reports{})
	db.DropTable(&model.StudentHistory{})
	db.DropTable(&model.Attends{})
	db.DropTable(&model.Enrollment{})
	db.DropTable(&model.Section{})

	db.DropTable(&model.Location{})
	db.DropTable(&model.Room{})
	db.DropTable(&model.Building{})

	db.DropTable(&model.TimeSlot{})
	db.DropTable(&model.Period{})
	db.DropTable(&model.Semester{})
	db.DropTable(&model.Day{})

	db.DropTable(&model.Prerequisite{})
	db.DropTable(&model.Course{})


	db.DropTable(&model.StudentMajor{})
	db.DropTable(&model.StudentMinor{})
	db.DropTable(&model.Minor{})
	db.DropTable(&model.Major{})

	db.DropTable(&model.FullTimeStudent{})
	db.DropTable(&model.PartTimeStudent{})
	db.DropTable(&model.Student{})

	db.DropTable(&model.PartTimeFaculty{})
	db.DropTable(&model.FullTimeFaculty{})
	db.DropTable(&model.Faculty{})

	db.DropTable(&model.Researcher{})
	db.DropTable(&model.Admin{})
	db.DropTable(&model.MainUser{})
	db.DropTable(&model.Department{})


	db.AutoMigrate(
		&model.MainUser{},
		&model.Student{},
		&model.PartTimeStudent{},
		&model.FullTimeStudent{},
		&model.Department{},
		&model.Faculty{},
		&model.PartTimeFaculty{},
		&model.FullTimeFaculty{},
		&model.Admin{},
		&model.Researcher{},
		&model.Major{},
		&model.Minor{},
		&model.StudentMajor{},
		&model.StudentMinor{},

		&model.Course{},
		&model.Prerequisite{},

		&model.Day{},
		&model.Semester{},
		&model.Period{},
		&model.TimeSlot{},

		&model.Building{},
		&model.Room{},
		&model.Location{},

		&model.Section{},
		&model.Enrollment{},
		&model.Attends{},
		&model.StudentHistory{},

		&model.Reports{},
	)

	db.Model(&model.Student{}).AddForeignKey("student_id", "main_user(user_id)", "CASCADE", "CASCADE")

	db.Model(&model.PartTimeStudent{}).AddForeignKey("part_time_student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.FullTimeStudent{}).AddForeignKey("full_time_student_id", "student(student_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Faculty{}).AddForeignKey("faculty_id", "main_user(user_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Faculty{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.PartTimeFaculty{}).AddForeignKey("part_time_faculty_id", "faculty(faculty_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.FullTimeFaculty{}).AddForeignKey("full_time_faculty_id", "faculty(faculty_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Admin{}).AddForeignKey("admin_id", "main_user(user_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Researcher{}).AddForeignKey("researcher_id", "main_user(user_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Major{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Minor{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.StudentMajor{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.StudentMajor{}).AddForeignKey("major_id", "major(major_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.StudentMinor{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.StudentMinor{}).AddForeignKey("minor_id", "major(major_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.TimeSlot{}).AddForeignKey("day_id", "day(day_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.TimeSlot{}).AddForeignKey("semester_id", "semester(semester_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.TimeSlot{}).AddForeignKey("period_id", "period(period_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Course{}).AddForeignKey("department_id", "department(department_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Prerequisite{}).AddForeignKey("course_required_by", "course(course_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Prerequisite{}).AddForeignKey("course_requirement", "course(course_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.TimeSlot{}).AddForeignKey("period_id", "period(period_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.TimeSlot{}).AddForeignKey("day_id", "day(day_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.TimeSlot{}).AddForeignKey("semester_id", "semester(semester_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Location{}).AddForeignKey("building_id", "building(building_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Location{}).AddForeignKey("room_id", "room(room_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Section{}).AddForeignKey("faculty_id", "faculty(faculty_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Section{}).AddForeignKey("time_slot_id", "time_slot(time_slot_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Section{}).AddForeignKey("course_id", "course(course_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Section{}).AddForeignKey("location_id", "location(location_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Enrollment{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Enrollment{}).AddForeignKey("section_id", "section(section_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Attends{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Attends{}).AddForeignKey("enrollment_id", "enrollment(enrollment_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.StudentHistory{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.StudentHistory{}).AddForeignKey("enrollment_id", "enrollment(enrollment_id)", "RESTRICT", "RESTRICT")

	db.Model(&model.Reports{}).AddForeignKey("researcher_id", "researcher(researcher_id)", "RESTRICT", "RESTRICT")


	user1 := model.MainUser{FirstName: "Pat", LastName:"Lagat", UserEmail:"plagat@yahoo.com", UserPassword:"pl12345", UserType:1}
	user2 := model.MainUser{FirstName: "Irish", LastName:"James", UserEmail:"jirish@yahoo.com", UserPassword:"ij12345", UserType:1}
	user3 := model.MainUser{FirstName: "Trey", LastName:"Gorkin", UserEmail:"gork@yahoo.com", UserPassword:"tg12345", UserType:1}
	user4 := model.MainUser{FirstName: "Testy", LastName:"McTest", UserEmail:"test@test.test", UserPassword:"testPW", UserType:1}
	user5 := model.MainUser{FirstName: "Faculty", LastName:"McFaculton", UserEmail:"faculty@test.test", UserPassword:"testPW", UserType:2}
	user6 := model.MainUser{FirstName: "Ima", LastName:"Faculty", UserEmail:"ifaculty@starfleet.edu", UserPassword:"testPW", UserType:2}
	user7 := model.MainUser{FirstName: "Jordi", LastName:"LaForge", UserEmail:"laforge@starfleet.edu", UserPassword:"testPW", UserType:2}
	user8 := model.MainUser{FirstName: "Admin", LastName:"McAdminton", UserEmail:"admin@test.test", UserPassword:"testPW", UserType:3}
	user9 := model.MainUser{FirstName: "George", LastName:"Admintonson", UserEmail:"adminson@starfleet.edu", UserPassword:"testPW", UserType:3}
	user10 := model.MainUser{FirstName: "Lesdo", LastName:"SomeResearch", UserEmail:"research@starfleet.edu", UserPassword:"testPW", UserType:4}

	db.Create(&user1)
	db.Create(&user2)
	db.Create(&user3)
	db.Create(&user4)
	db.Create(&user5)
	db.Create(&user6)
	db.Create(&user7)
	db.Create(&user8)
	db.Create(&user9)
	db.Create(&user10)

	student2 := model.Student{StudentID: user2.UserID, StudentType:1}
	student1 := model.Student{StudentID: user1.UserID, StudentType:1}
	student3 := model.Student{StudentID: user3.UserID, StudentType:2}
	student4 := model.Student{StudentID: user4.UserID, StudentType:2}
	db.Create(&student1)
	db.Create(&student2)
	db.Create(&student3)
	db.Create(&student4)

	// example of finding related models
	u := model.MainUser{}
	db.Model(&student2).Association("MainUser").Find(&u)
	fmt.Println("For the student2, the user email is: ", u.UserEmail)

	fullTimeStudent1 := model.FullTimeStudent{FullTimeStudentID: student1.StudentID, NumCredits:16}
	fullTimeStudent2 := model.FullTimeStudent{FullTimeStudentID: student2.StudentID, NumCredits:18}
	db.Create(&fullTimeStudent1)
	db.Create(&fullTimeStudent2)
	partTimeStudent1 := model.PartTimeStudent{PartTimeStudentID: student3.StudentID, NumCredits:12}
	partTimeStudent2 := model.PartTimeStudent{PartTimeStudentID: student4.StudentID, NumCredits:8}
	db.Create(&partTimeStudent1)
	db.Create(&partTimeStudent2)

	// example of lookup back from part_time_student to main_user
	uLookup := model.MainUser{}
	stuLookup := model.Student{}
	db.Model(&partTimeStudent1).Association("Student").Find(&stuLookup)
	db.Model(&stuLookup).Association("MainUser").Find(&uLookup)
	fmt.Println("For the partTimeStudent1, the user email is: ", uLookup.UserEmail)

	department1 := model.Department{DepartmentName:"Math", DepartmentBuilding:"MathBuilding", DepartmentRoomNumber:"302", DepartmentChair:"Mr Math", DepartmentPhoneNumber:"111-222-3333"}
	department2 := model.Department{DepartmentName:"Computer Science", DepartmentBuilding:"CSBuilding", DepartmentRoomNumber:"100", DepartmentChair:"Mr Computer", DepartmentPhoneNumber:"123-456-7899"}
	db.Create(&department1)
	db.Create(&department2)

	faculty1 := model.Faculty{FacultyID:user5.UserID, FacultyType:1, DepartmentID:department1.DepartmentID}
	faculty2 := model.Faculty{FacultyID:user6.UserID, FacultyType:2, DepartmentID:department2.DepartmentID}
	faculty3 := model.Faculty{FacultyID:user7.UserID, FacultyType:1, DepartmentID:department1.DepartmentID}
	db.Create(&faculty1)
	db.Create(&faculty2)
	db.Create(&faculty3)

	// example of finding a many-one lookup
	dep := model.Department{}
	db.Model(&faculty1).Association("Department").Find(&dep)
	fmt.Println("For the faculty1, the department is: ", dep.DepartmentName)
	facMembers := []model.Faculty{}
	// search faculty by department (.Association and .Related doesn't seem to work)
	//db.Model(&department1).Association("Faculty").Find(&facMembers)
	db.Where(model.Faculty{DepartmentID: department1.DepartmentID}).Find(&facMembers)
	fmt.Println("For the departent1 (Math), the faculty is: ")
	for _, v := range facMembers {
		fmt.Println("FacultyID", v.FacultyID)
	}


	admin1 := model.Admin{AdminID:user8.UserID}
	admin2 := model.Admin{AdminID:user9.UserID}
	db.Create(&admin1)
	db.Create(&admin2)

	researcher1 := model.Researcher{ResearcherID:user10.UserID}
	db.Create(researcher1)


	major1 := model.Major{DepartmentID: department1.DepartmentID, MajorName:"MATH"}
	major2 := model.Major{DepartmentID: department1.DepartmentID, MajorName:"SUPERMATH"}
	major3 := model.Major{DepartmentID: department2.DepartmentID, MajorName:"CIS"}
	major4 := model.Major{DepartmentID: department2.DepartmentID, MajorName:"MIS"}
	db.Create(&major1)
	db.Create(&major2)
	db.Create(&major3)
	db.Create(&major4)

	minor1 := model.Minor{DepartmentID: department1.DepartmentID, MinorName:"APPLIED MATH"}
	minor2 := model.Minor{DepartmentID: department1.DepartmentID, MinorName:"MINOR MATH"}
	minor3 := model.Minor{DepartmentID: department2.DepartmentID, MinorName:"LITTLE COMPUTERS"}
	minor4 := model.Minor{DepartmentID: department2.DepartmentID, MinorName:"COMPUTER STUFF"}
	db.Create(&minor1)
	db.Create(&minor2)
	db.Create(&minor3)
	db.Create(&minor4)

	course1 := model.Course{CourseName:"Warp Field Mechanics", CourseCredits:4, DepartmentID:department1.DepartmentID,
		CourseDescription:"An introduction to the theory behind faster than light space travel."}

	course2 := model.Course{CourseName:"History of Space Travel", CourseCredits:4, DepartmentID:department1.DepartmentID,
		CourseDescription:"A survey of humanity's entry into the age of space exploration, from the first artificial satellite to first contact with the Vulcans. "}

	course3 := model.Course{CourseName:"Contemporary Holography", CourseCredits:4, DepartmentID:department1.DepartmentID,
		CourseDescription:"An introduction to Holography."}

	course4 := model.Course{CourseName:"Newtonian Physics I", CourseCredits:4, DepartmentID:department2.DepartmentID, CourseDescription:"The basic building blocks of Newtonain physics icluding kinetics, force, rotation, and harmonic motion"}
	course5 := model.Course{CourseName:"Newtonian Physics II", CourseCredits:4, DepartmentID:department2.DepartmentID, CourseDescription:"The fundamentals of optics, electricity and magnetism"}

	db.Create(&course1)
	db.Create(&course2)
	db.Create(&course3)
	db.Create(&course4)
	db.Create(&course5)

	preReq1 := model.Prerequisite{CourseRequiredBy:course5.CourseID, CourseRequirement: course4.CourseID}
	preReq2 := model.Prerequisite{CourseRequiredBy:course1.CourseID, CourseRequirement: course2.CourseID}
	db.Create(&preReq1)
	db.Create(&preReq2)

	building := model.Building{BuildingName:"TheBuilding", BuildingAddress:"5 Shady Grove"}
	db.Create(&building)

	room1 := model.Room{RoomNumber:"B100", RoomType:"Lecture"}
	room2 := model.Room{RoomNumber:"C200", RoomType: "LAB"}
	db.Create(&room1)
	db.Create(&room2)

	location1 := model.Location{BuildingID:building.BuildingID, RoomID:room1.RoomID}
	location2 := model.Location{BuildingID:building.BuildingID, RoomID:room2.RoomID}
	db.Create(&location1)
	db.Create(&location2)

	day1 := model.Day{MeetingDay:"MW"}
	day2 := model.Day{MeetingDay:"TR"}
	db.Create(&day1)
	db.Create(&day2)

	semester1 := model.Semester{Year:2017, Season:"Spring"}
	semester2 := model.Semester{Year: 2018, Season:"Fall"}
	db.Create(&semester1)
	db.Create(&semester2)

	timeform := "Jan 2, 2006 at 3:04pm (MST)"
	t1, _ := time.Parse(timeform, "Jan 2, 2006 at 1:00pm (MST)")
	t2, _ := time.Parse(timeform, "Jan 2, 2006 at 2:30pm (MST)")
	t3, _ := time.Parse(timeform, "Jan 2, 2006 at 10:00pm (MST)")
	t4, _ := time.Parse(timeform, "Jan 2, 2006 at 11:30pm (MST)")
	period1 := model.Period{StartTime: t1, EndTime: t2}
	period2 := model.Period{StartTime: t3, EndTime: t4}
	db.Create(&period1)
	db.Create(&period2)

	timeslot1 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day1.DayID, PeriodID:period1.PeriodID}
	timeslot2 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day1.DayID, PeriodID:period2.PeriodID}
	timeslot3 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period1.PeriodID}
	timeslot4 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period2.PeriodID}
	db.Create(&timeslot1)
	db.Create(&timeslot2)
	db.Create(&timeslot3)
	db.Create(&timeslot4)

	section1 := model.Section{CourseSectionNumber:001, CourseID:course1.CourseID, FacultyID:faculty1.FacultyID, TimeSlotID:timeslot1.TimeSlotID, LocationID:location1.LocationID}
	section2 := model.Section{CourseSectionNumber:002, CourseID:course1.CourseID, FacultyID:faculty1.FacultyID, TimeSlotID:timeslot2.TimeSlotID, LocationID:location1.LocationID}
	section3 := model.Section{CourseSectionNumber:001, CourseID:course2.CourseID, FacultyID:faculty2.FacultyID, TimeSlotID:timeslot1.TimeSlotID, LocationID:location2.LocationID}
	section4 := model.Section{CourseSectionNumber:002, CourseID:course2.CourseID, FacultyID:faculty2.FacultyID, TimeSlotID:timeslot2.TimeSlotID, LocationID:location2.LocationID}
	section5 := model.Section{CourseSectionNumber:001, CourseID:course3.CourseID, FacultyID:faculty1.FacultyID, TimeSlotID:timeslot3.TimeSlotID, LocationID:location1.LocationID}
	section6 := model.Section{CourseSectionNumber:001, CourseID:course4.CourseID, FacultyID:faculty3.FacultyID, TimeSlotID:timeslot3.TimeSlotID, LocationID:location2.LocationID}
	section7 := model.Section{CourseSectionNumber:002, CourseID:course4.CourseID, FacultyID:faculty1.FacultyID, TimeSlotID:timeslot4.TimeSlotID, LocationID:location1.LocationID}
	section8 := model.Section{CourseSectionNumber:001, CourseID:course5.CourseID, FacultyID:faculty3.FacultyID, TimeSlotID:timeslot4.TimeSlotID, LocationID:location2.LocationID}
	db.Create(&section1)
	db.Create(&section2)
	db.Create(&section3)
	db.Create(&section4)
	db.Create(&section5)
	db.Create(&section6)
	db.Create(&section7)
	db.Create(&section8)

	enroll1 := model.Enrollment{StudentID:student1.StudentID, SectionID:section3.SectionID, Grade:"-"}
	enroll2 := model.Enrollment{StudentID:student2.StudentID, SectionID:section3.SectionID, Grade:"-"}
	enroll3 := model.Enrollment{StudentID:student3.StudentID, SectionID:section4.SectionID, Grade:"-"}
	enroll4 := model.Enrollment{StudentID:student4.StudentID, SectionID:section4.SectionID, Grade:"-"}
	enroll5 := model.Enrollment{StudentID:student1.StudentID, SectionID:section6.SectionID, Grade:"-"}
	enroll6 := model.Enrollment{StudentID:student2.StudentID, SectionID:section6.SectionID, Grade:"-"}
	enroll7 := model.Enrollment{StudentID:student3.StudentID, SectionID:section6.SectionID, Grade:"-"}
	enroll8 := model.Enrollment{StudentID:student4.StudentID, SectionID:section7.SectionID, Grade:"-"}
	enroll9 := model.Enrollment{StudentID:student1.StudentID, SectionID:section5.SectionID, Grade:"-"}
	enroll10 := model.Enrollment{StudentID:student2.StudentID, SectionID:section5.SectionID, Grade:"-"}
	enroll11 := model.Enrollment{StudentID:student3.StudentID, SectionID:section1.SectionID, Grade:"-"}
	enroll12 := model.Enrollment{StudentID:student1.StudentID, SectionID:section2.SectionID, Grade:"-"}
	db.Create(&enroll1)
	db.Create(&enroll2)
	db.Create(&enroll3)
	db.Create(&enroll4)
	db.Create(&enroll5)
	db.Create(&enroll6)
	db.Create(&enroll7)
	db.Create(&enroll8)
	db.Create(&enroll9)
	db.Create(&enroll10)
	db.Create(&enroll11)
	db.Create(&enroll12)

	history1 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll1.EnrollmentID, Status:"In progress"}
	history2 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll2.EnrollmentID, Status:"In progress"}
	history3 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll3.EnrollmentID, Status:"In progress"}
	history4 := model.StudentHistory{StudentID:student4.StudentID, EnrollmentID:enroll4.EnrollmentID, Status:"In progress"}
	history5 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll5.EnrollmentID, Status:"In progress"}
	history6 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll6.EnrollmentID, Status:"In progress"}
	history7 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll7.EnrollmentID, Status:"In progress"}
	history8 := model.StudentHistory{StudentID:student4.StudentID, EnrollmentID:enroll8.EnrollmentID, Status:"In progress"}
	history9 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll9.EnrollmentID, Status:"In progress"}
	history10 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll10.EnrollmentID, Status:"In progress"}
	history11 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll11.EnrollmentID, Status:"In progress"}
	history12 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll12.EnrollmentID, Status:"In progress"}
	db.Create(&history1)
	db.Create(&history2)
	db.Create(&history3)
	db.Create(&history4)
	db.Create(&history5)
	db.Create(&history6)
	db.Create(&history7)
	db.Create(&history8)
	db.Create(&history9)
	db.Create(&history10)
	db.Create(&history11)
	db.Create(&history12)
}


*/

