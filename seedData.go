package main


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"Starfleet/model"
	"os"
	"fmt"
	//"time"
)
func main() {
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

	db.DropTable(&model.StudentHolds{})
	db.DropTable(&model.Hold{})

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

		&model.Hold{},
		&model.StudentHolds{},

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

	user1 := model.MainUser{FirstName: "Pat", LastName: "Lagat", UserEmail: "plagat@yahoo.com", UserPassword: "pl12345", UserType: 1}
	user2 := model.MainUser{FirstName: "Irish", LastName: "James", UserEmail: "jirish@yahoo.com", UserPassword: "ij12345", UserType: 1}
	user3 := model.MainUser{FirstName: "Trey", LastName: "Gorkin", UserEmail: "gork@yahoo.com", UserPassword: "tg12345", UserType: 1}
	user4 := model.MainUser{FirstName: "Testy", LastName: "McTest", UserEmail: "test@test.test", UserPassword: "testPW", UserType: 1}
	user5 := model.MainUser{FirstName: "Faculty", LastName: "McFaculton", UserEmail: "faculty@test.test", UserPassword: "testPW", UserType: 2}
	user6 := model.MainUser{FirstName: "Aymen", LastName: "Johnson", UserEmail: "ifaculty@starfleet.edu", UserPassword: "testPW", UserType: 2}
	user7 := model.MainUser{FirstName: "Jordi", LastName: "LaForge", UserEmail: "laforge@starfleet.edu", UserPassword: "testPW", UserType: 2}
	user8 := model.MainUser{FirstName: "Admin", LastName: "McAdminton", UserEmail: "admin@test.test", UserPassword: "testPW", UserType: 3}
	user9 := model.MainUser{FirstName: "George", LastName: "Admintonson", UserEmail: "adminson@starfleet.edu", UserPassword: "testPW", UserType: 3}
	user10 := model.MainUser{FirstName: "Lesdo", LastName: "SomeResearch", UserEmail: "research@starfleet.edu", UserPassword: "testPW", UserType: 4}
	facultyUser4 := model.MainUser{FirstName: "Brandon", LastName: "Sanderson", UserEmail: "sanderson@gmail.com", UserPassword: "testPW", UserType: 2}
	facultyUser5 := model.MainUser{FirstName: "Mike", LastName: "Tyson", UserEmail: "tyson@gmail.com", UserPassword: "testPW", UserType: 2}
	facultyUser6 := model.MainUser{FirstName: "Eddard", LastName: "Stark", UserEmail: "estark@winterfell.com", UserPassword: "testPW", UserType: 2}

	fmt.Println("Creating users")
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
	db.Create(&facultyUser4)
	db.Create(&facultyUser5)
	db.Create(&facultyUser6)

	student2 := model.Student{StudentID: user2.UserID, StudentType: 1}
	student1 := model.Student{StudentID: user1.UserID, StudentType: 1}
	student3 := model.Student{StudentID: user3.UserID, StudentType: 2}
	student4 := model.Student{StudentID: user4.UserID, StudentType: 2}
	db.Create(&student1)
	db.Create(&student2)
	db.Create(&student3)
	db.Create(&student4)

	// example of finding related models
	u := model.MainUser{}
	db.Model(&student2).Association("MainUser").Find(&u)
	fmt.Println("For the student2, the user email is: ", u.UserEmail)

	fullTimeStudent1 := model.FullTimeStudent{FullTimeStudentID: student1.StudentID, NumCredits: 16}
	fullTimeStudent2 := model.FullTimeStudent{FullTimeStudentID: student2.StudentID, NumCredits: 18}
	db.Create(&fullTimeStudent1)
	db.Create(&fullTimeStudent2)
	partTimeStudent1 := model.PartTimeStudent{PartTimeStudentID: student3.StudentID, NumCredits: 12}
	partTimeStudent2 := model.PartTimeStudent{PartTimeStudentID: student4.StudentID, NumCredits: 8}
	db.Create(&partTimeStudent1)
	db.Create(&partTimeStudent2)

	// example of lookup back from part_time_student to main_user
	uLookup := model.MainUser{}
	stuLookup := model.Student{}
	db.Model(&partTimeStudent1).Association("Student").Find(&stuLookup)
	db.Model(&stuLookup).Association("MainUser").Find(&uLookup)
	fmt.Println("For the partTimeStudent1, the user email is: ", uLookup.UserEmail)

	department1 := model.Department{DepartmentName: "Math", DepartmentBuilding: "MathBuilding", DepartmentRoomNumber: "302", DepartmentChair: "Techanie, Geta", DepartmentPhoneNumber: "111-222-3333"}
	department2 := model.Department{DepartmentName: "Computer Science", DepartmentBuilding: "CSBuilding", DepartmentRoomNumber: "100", DepartmentChair: "Skiena, Steven", DepartmentPhoneNumber: "123-456-7899"}
	department3 := model.Department{DepartmentName: "History", DepartmentBuilding: "HistoryBuilding", DepartmentRoomNumber: "200", DepartmentChair: "Smith, Jackson", DepartmentPhoneNumber: "516-553-3291"}
	department4 := model.Department{DepartmentName: "Liberal Arts", DepartmentBuilding: "ArtsBuilding", DepartmentRoomNumber: "200", DepartmentChair: "Johnson, Allan", DepartmentPhoneNumber: "516-553-3291"}
	department5 := model.Department{DepartmentName: "Business", DepartmentBuilding: "BusinessBuilding", DepartmentRoomNumber: "400", DepartmentChair: "Seymour, Roswalt", DepartmentPhoneNumber: "516-553-3291"}

	fmt.Println("Creating departments")

	db.Create(&department1)
	db.Create(&department2)
	db.Create(&department3)
	db.Create(&department4)
	db.Create(&department5)

	faculty1 := model.Faculty{FacultyID: user5.UserID, FacultyType: 1, DepartmentID: department1.DepartmentID}
	faculty2 := model.Faculty{FacultyID: user6.UserID, FacultyType: 2, DepartmentID: department2.DepartmentID}
	faculty3 := model.Faculty{FacultyID: user7.UserID, FacultyType: 1, DepartmentID: department1.DepartmentID}

	faculty4 := model.Faculty{FacultyID: facultyUser4.UserID, FacultyType: 1, DepartmentID: department3.DepartmentID}
	faculty5 := model.Faculty{FacultyID: facultyUser5.UserID, FacultyType: 1, DepartmentID: department4.DepartmentID}
	faculty6 := model.Faculty{FacultyID: facultyUser6.UserID, FacultyType: 1, DepartmentID: department5.DepartmentID}


	db.Create(&faculty1)
	db.Create(&faculty2)
	db.Create(&faculty3)
	db.Create(&faculty4)
	db.Create(&faculty5)
	db.Create(&faculty6)

	// example of finding a many-one lookup
	dep := model.Department{}
	db.Model(&faculty1).Association("Department").Find(&dep)
	fmt.Println("For the faculty1, the department is: ", dep.DepartmentName)
	facMembers := []model.Faculty{}
	// search faculty by department (.Association and .Related doesn't seem to work)
	//db.Model(&department1).Association("Faculty").Find(&facMembers)
	db.Where(model.Faculty{DepartmentID: department1.DepartmentID}).Find(&facMembers)
	fmt.Println("For the departent1 (Math), the faculty is: ")

	admin1 := model.Admin{AdminID: user8.UserID}
	admin2 := model.Admin{AdminID: user9.UserID}
	db.Create(&admin1)
	db.Create(&admin2)

	researcher1 := model.Researcher{ResearcherID: user10.UserID}
	db.Create(&researcher1)

	major1 := model.Major{DepartmentID: department1.DepartmentID, MajorName: "MATH"}
	major2 := model.Major{DepartmentID: department1.DepartmentID, MajorName: "SUPERMATH"}
	major3 := model.Major{DepartmentID: department2.DepartmentID, MajorName: "CIS"}
	major4 := model.Major{DepartmentID: department2.DepartmentID, MajorName: "MIS"}
	db.Create(&major1)
	db.Create(&major2)
	db.Create(&major3)
	db.Create(&major4)

	minor1 := model.Minor{DepartmentID: department1.DepartmentID, MinorName: "APPLIED MATH"}
	minor2 := model.Minor{DepartmentID: department1.DepartmentID, MinorName: "MINOR MATH"}
	minor3 := model.Minor{DepartmentID: department2.DepartmentID, MinorName: "LITTLE COMPUTERS"}
	minor4 := model.Minor{DepartmentID: department2.DepartmentID, MinorName: "COMPUTER STUFF"}
	db.Create(&minor1)
	db.Create(&minor2)
	db.Create(&minor3)
	db.Create(&minor4)

	hold1 := model.Hold{HoldName: "Unpaid Bill"}
	hold2 := model.Hold{HoldName: "Un-submitted Health Forms"}
	hold3 := model.Hold{HoldName: "Unpaid Parking Ticket"}
	hold4 := model.Hold{HoldName: "Unpaid Speeding Ticket"}

	db.Create(&hold1)
	db.Create(&hold2)
	db.Create(&hold3)
	db.Create(&hold4)

	studenthold1 := model.StudentHolds{StudentID: student1.StudentID, HoldID: hold1.HoldID}
	studenthold2 := model.StudentHolds{StudentID: student1.StudentID, HoldID: hold2.HoldID}
	studenthold3 := model.StudentHolds{StudentID: student2.StudentID, HoldID: hold3.HoldID}
	studenthold4 := model.StudentHolds{StudentID: student1.StudentID, HoldID: hold4.HoldID}

	db.Create(&studenthold1)
	db.Create(&studenthold2)
	db.Create(&studenthold3)
	db.Create(&studenthold4)

	course1 := model.Course{CourseName: "Warp Field Mechanics", CourseCredits: 4, DepartmentID: department1.DepartmentID,
		CourseDescription:          "An introduction to the theory behind faster than light space travel."}

	course2 := model.Course{CourseName: "History of Space Travel", CourseCredits: 4, DepartmentID: department1.DepartmentID,
		CourseDescription:          "A survey of humanity's entry into the age of space exploration, from the first artificial satellite to first contact with the Vulcans. "}

	course3 := model.Course{CourseName: "Contemporary Holography", CourseCredits: 4, DepartmentID: department1.DepartmentID,
		CourseDescription:          "An introduction to Holography."}

	course4 := model.Course{CourseName: "Newtonian Physics I", CourseCredits: 4, DepartmentID: department2.DepartmentID, CourseDescription: "The basic building blocks of Newtonain physics icluding kinetics, force, rotation, and harmonic motion"}
	course5 := model.Course{CourseName: "Newtonian Physics II", CourseCredits: 4, DepartmentID: department2.DepartmentID, CourseDescription: "The fundamentals of optics, electricity and magnetism"}
	db.Create(&course1)
	db.Create(&course2)
	db.Create(&course3)
	db.Create(&course4)
	db.Create(&course5)

	preReq1 := model.Prerequisite{CourseRequiredBy: course5.CourseID, CourseRequirement: course4.CourseID}
	preReq2 := model.Prerequisite{CourseRequiredBy: course1.CourseID, CourseRequirement: course2.CourseID}
	db.Create(&preReq1)
	db.Create(&preReq2)

	building := model.Building{BuildingName: "The Academy", BuildingAddress: "5 Shawsington Road"}
	building2 := model.Building{BuildingName: "Riften Building", BuildingAddress: "115 Shawsington Road"}

	db.Create(&building)
	db.Create(&building2)

	//building 1
	room1 := model.Room{RoomNumber: "B100", RoomType: "Lecture Hall", RoomCapacity: 100}
	room2 := model.Room{RoomNumber: "C200", RoomType: "LAB", RoomCapacity: 10}
	room3 := model.Room{RoomNumber: "C210", RoomType: "Classroom", RoomCapacity: 30}
	room4 := model.Room{RoomNumber: "C220", RoomType: "Classroom", RoomCapacity: 30}
	room5 := model.Room{RoomNumber: "C230", RoomType: "Classroom", RoomCapacity: 40}
	room6 := model.Room{RoomNumber: "C240", RoomType: "Classroom", RoomCapacity: 25}

	//building 2
	b2room1 := model.Room{RoomNumber: "P100", RoomType: "Classroom", RoomCapacity: 25}
	b2room2 := model.Room{RoomNumber: "P150", RoomType: "Classroom", RoomCapacity: 25}
	b2room3 := model.Room{RoomNumber: "P180", RoomType: "Classroom", RoomCapacity: 30}
	b2room4 := model.Room{RoomNumber: "H100", RoomType: "Lecture Hall", RoomCapacity: 125}

	db.Create(&b2room1)
	db.Create(&b2room2)
	db.Create(&b2room3)
	db.Create(&b2room4)

	db.Create(&room1)
	db.Create(&room2)
	db.Create(&room3)
	db.Create(&room4)
	db.Create(&room5)
	db.Create(&room6)

	fmt.Println("Creating locations")


	location1 := model.Location{BuildingID: building.BuildingID, RoomID: room1.RoomID}
	location2 := model.Location{BuildingID: building.BuildingID, RoomID: room2.RoomID}
	location3 := model.Location{BuildingID: building.BuildingID, RoomID: room3.RoomID}
	location4 := model.Location{BuildingID: building.BuildingID, RoomID: room4.RoomID}
	location5 := model.Location{BuildingID: building.BuildingID, RoomID: room5.RoomID}
	location6 := model.Location{BuildingID: building.BuildingID, RoomID: room6.RoomID}

	location7 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room1.RoomID}
	location8 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room2.RoomID}
	location9 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room3.RoomID}
	location10 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room4.RoomID}

	db.Create(&location1)
	db.Create(&location2)
	db.Create(&location3)
	db.Create(&location4)
	db.Create(&location5)
	db.Create(&location6)

	db.Create(&location7)
	db.Create(&location8)
	db.Create(&location9)
	db.Create(&location10)

	day1 := model.Day{MeetingDay: "MW"}
	day2 := model.Day{MeetingDay: "TR"}
	db.Create(&day1)
	db.Create(&day2)

	semester1 := model.Semester{Year: 2018, Season: "Spring"}
	semester2 := model.Semester{Year: 2018, Season: "Fall"}
	winterSemester := model.Semester{Year:2018, Season:"Winter"}
	summerSemester := model.Semester{Year:2018, Season:"Summer"}

	fallLastYear := model.Semester{Year:2017, Season:"Fall"}
	springLastYear := model.Semester{Year:2017, Season:"Spring"}
	winterLastYear := model.Semester{Year:2017, Season:"Winter"}
	summerLastYear := model.Semester{Year:2017, Season:"Summer"}

	fallLastLastYear := model.Semester{Year:2016, Season:"Fall"}
	springLastLastYear := model.Semester{Year:2016, Season:"Spring"}

	fallLastLastLastYear := model.Semester{Year:2015, Season:"Fall"}
	springLastLastLastYear := model.Semester{Year:2015, Season:"Spring"}

	fmt.Println("Creating semesters")

	db.Create(&semester1)
	db.Create(&semester2)
	db.Create(&winterSemester)
	db.Create(&summerSemester)
	db.Create(&fallLastYear)
	db.Create(&springLastYear)

	db.Create(&winterLastYear)
	db.Create(&summerLastYear)
	db.Create(&fallLastLastYear)
	db.Create(&springLastLastYear)
	db.Create(&fallLastLastLastYear)
	db.Create(&springLastLastLastYear)

	/*
	timeform := "Jan 2, 2006 at 3:04pm (MST)"
	t1, _ := time.Parse(timeform, "Jan 2, 2006 at 1:00pm (MST)")
	t2, _ := time.Parse(timeform, "Jan 2, 2006 at 2:30pm (MST)")
	t3, _ := time.Parse(timeform, "Jan 2, 2006 at 10:00pm (MST)")
	t4, _ := time.Parse(timeform, "Jan 2, 2006 at 11:30pm (MST)")
	period1 := model.Period{Star	tTime: t1, EndTime: t2}
	period2 := model.Period{StartTime: t3, EndTime: t4}
	*/


	period0 := model.Period{Time:"9:40 AM - 11:10 AM"}
	period1 := model.Period{Time:"11:20 AM - 12:50 PM"}
	period2 := model.Period{Time:"3:50 PM - 5:20 PM"}
	period3 := model.Period{Time:"5:30 PM - 7:00 PM"}
	period4 := model.Period{Time:"7:10 PM - 8:40 PM"}
	db.Create(&period0)
	db.Create(&period1)
	db.Create(&period2)
	db.Create(&period3)
	db.Create(&period4)

	//Timeslot for spring 2018
		//all 5 periods for MW
	timeslot1 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day1.DayID, PeriodID:period1.PeriodID}
	timeslot2 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day1.DayID, PeriodID:period2.PeriodID}
	timeslota := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day1.DayID, PeriodID:period3.PeriodID}
	timeslotb := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day1.DayID, PeriodID:period4.PeriodID}
	timeslotc := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day1.DayID, PeriodID:period0.PeriodID}

		//all 5 periods for TR
	timeslot11 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period1.PeriodID}
	timeslot21 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period2.PeriodID}
	timeslota1 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period3.PeriodID}
	timeslotb2 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period4.PeriodID}
	timeslotc3 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period0.PeriodID}

	timeslot3 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period1.PeriodID}
	timeslot4 := model.TimeSlot{SemesterID:semester1.SemesterID, DayID:day2.DayID, PeriodID:period2.PeriodID}


	//Timeslot for Fall 2017
		//MW
	timeslotFall2017 := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day1.DayID, PeriodID:period1.PeriodID}
	timeslotFall12017 := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day1.DayID, PeriodID:period2.PeriodID}
	timeslotFall22017 := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day1.DayID, PeriodID:period3.PeriodID}
	timeslotFall32017 := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day1.DayID, PeriodID:period4.PeriodID}
	timeslotFall42017 := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day1.DayID, PeriodID:period0.PeriodID}

	db.Create(&timeslotFall2017)
	db.Create(&timeslotFall12017)
	db.Create(&timeslotFall22017)
	db.Create(&timeslotFall32017)
	db.Create(&timeslotFall42017)


		//TR
	timeslotFall2017tr := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day2.DayID, PeriodID:period1.PeriodID}
	timeslotFall12017tr := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day2.DayID, PeriodID:period2.PeriodID}
	timeslotFall22017tr := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day2.DayID, PeriodID:period3.PeriodID}
	timeslotFall32017tr := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day2.DayID, PeriodID:period4.PeriodID}
	timeslotFall42017tr := model.TimeSlot{SemesterID:fallLastYear.SemesterID, DayID:day2.DayID, PeriodID:period0.PeriodID}

	db.Create(&timeslotFall2017tr)
	db.Create(&timeslotFall12017tr)
	db.Create(&timeslotFall22017tr)
	db.Create(&timeslotFall32017tr)
	db.Create(&timeslotFall42017tr)

	// Timeslot for Spring 2017
		//MW
	timeslotSpring2017 := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day1.DayID, PeriodID:period1.PeriodID}
	timeslotSpring12017 := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day1.DayID, PeriodID:period2.PeriodID}
	timeslotSpring22017 := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day1.DayID, PeriodID:period3.PeriodID}
	timeslotSpring32017 := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day1.DayID, PeriodID:period4.PeriodID}
	timeslotSpring42017 := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day1.DayID, PeriodID:period0.PeriodID}

	db.Create(&timeslotSpring2017)
	db.Create(&timeslotSpring12017)
	db.Create(&timeslotSpring22017)
	db.Create(&timeslotSpring32017)
	db.Create(&timeslotSpring42017)


		//TR
	timeslotSpring2017tr := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day2.DayID, PeriodID:period1.PeriodID}
	timeslotSpring12017tr := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day2.DayID, PeriodID:period2.PeriodID}
	timeslotSpring22017tr := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day2.DayID, PeriodID:period3.PeriodID}
	timeslotSpring32017tr := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day2.DayID, PeriodID:period4.PeriodID}
	timeslotSpring42017tr := model.TimeSlot{SemesterID:springLastYear.SemesterID, DayID:day2.DayID, PeriodID:period0.PeriodID}

	db.Create(&timeslotSpring2017tr)
	db.Create(&timeslotSpring12017tr)
	db.Create(&timeslotSpring22017tr)
	db.Create(&timeslotSpring32017tr)
	db.Create(&timeslotSpring42017tr)

	//Timeslot for Fall 2016
	//MW
	timeslotFall2016 := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day1.DayID, PeriodID:period1.PeriodID}
	timeslotFall12016 := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day1.DayID, PeriodID:period2.PeriodID}
	timeslotFall22016 := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day1.DayID, PeriodID:period3.PeriodID}
	timeslotFall32016 := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day1.DayID, PeriodID:period4.PeriodID}
	timeslotFall42016 := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day1.DayID, PeriodID:period0.PeriodID}

	db.Create(&timeslotFall2016)
	db.Create(&timeslotFall12016)
	db.Create(&timeslotFall22016)
	db.Create(&timeslotFall32016)
	db.Create(&timeslotFall42016)


	//TR
	timeslotFall2016tr := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day2.DayID, PeriodID:period1.PeriodID}
	timeslotFall12016tr := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day2.DayID, PeriodID:period2.PeriodID}
	timeslotFall22016tr := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day2.DayID, PeriodID:period3.PeriodID}
	timeslotFall32016tr := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day2.DayID, PeriodID:period4.PeriodID}
	timeslotFall42016tr := model.TimeSlot{SemesterID:fallLastLastYear.SemesterID, DayID:day2.DayID, PeriodID:period0.PeriodID}

	db.Create(&timeslotFall2016tr)
	db.Create(&timeslotFall12016tr)
	db.Create(&timeslotFall22016tr)
	db.Create(&timeslotFall32016tr)
	db.Create(&timeslotFall42016tr)

	//Timeslot for Spring 2015

	fmt.Println("Creating timeslots")

	db.Create(&timeslot1)
	db.Create(&timeslot2)
	db.Create(&timeslot3)
	db.Create(&timeslot4)
	db.Create(&timeslota)
	db.Create(&timeslotb)
	db.Create(&timeslotc)
	db.Create(&timeslot11)
	db.Create(&timeslot21)
	db.Create(&timeslota1)
	db.Create(&timeslotb2)
	db.Create(&timeslotc3)


	//Spring 2018 Sections //These are wrong all these enrollments are for spring 2018, should not be in progress until the next semester
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

	//Fall 2017 Sections
	sectionFall2017a := model.Section{CourseSectionNumber:001, CourseID:course1.CourseID, FacultyID:faculty1.FacultyID,TimeSlotID:timeslotFall2017.TimeSlotID, LocationID:location4.LocationID}
	sectionFall2017b := model.Section{CourseSectionNumber:002, CourseID:course2.CourseID, FacultyID:faculty2.FacultyID,TimeSlotID:timeslotFall12017.TimeSlotID, LocationID:location5.LocationID}
	sectionFall2017c := model.Section{CourseSectionNumber:001, CourseID:course3.CourseID, FacultyID:faculty3.FacultyID,TimeSlotID:timeslotFall22017.TimeSlotID, LocationID:location6.LocationID}
	sectionFall2017d := model.Section{CourseSectionNumber:002, CourseID:course4.CourseID, FacultyID:faculty4.FacultyID,TimeSlotID:timeslotFall32017.TimeSlotID, LocationID:location7.LocationID}
	sectionFall2017e := model.Section{CourseSectionNumber:001, CourseID:course5.CourseID, FacultyID:faculty5.FacultyID,TimeSlotID:timeslotFall42017.TimeSlotID, LocationID:location8.LocationID}
	sectionFall2017f := model.Section{CourseSectionNumber:003, CourseID:course1.CourseID, FacultyID:faculty6.FacultyID,TimeSlotID:timeslotFall2017tr.TimeSlotID, LocationID:location9.LocationID}
	sectionFall2017g := model.Section{CourseSectionNumber:002, CourseID:course2.CourseID, FacultyID:faculty1.FacultyID,TimeSlotID:timeslotFall12017tr.TimeSlotID, LocationID:location10.LocationID}
	sectionFall2017h := model.Section{CourseSectionNumber:001, CourseID:course3.CourseID, FacultyID:faculty2.FacultyID,TimeSlotID:timeslotFall22017tr.TimeSlotID, LocationID:location1.LocationID}
	sectionFall2017i := model.Section{CourseSectionNumber:002, CourseID:course4.CourseID, FacultyID:faculty3.FacultyID,TimeSlotID:timeslotFall32017tr.TimeSlotID, LocationID:location2.LocationID}
	sectionFall2017j := model.Section{CourseSectionNumber:003, CourseID:course5.CourseID, FacultyID:faculty4.FacultyID,TimeSlotID:timeslotFall42017tr.TimeSlotID, LocationID:location3.LocationID}
	sectionFall2017k := model.Section{CourseSectionNumber:001, CourseID:course2.CourseID, FacultyID:faculty5.FacultyID,TimeSlotID:timeslotFall2017tr.TimeSlotID, LocationID:location4.LocationID}

	db.Create(&sectionFall2017a)
	db.Create(&sectionFall2017b)
	db.Create(&sectionFall2017c)
	db.Create(&sectionFall2017d)
	db.Create(&sectionFall2017e)
	db.Create(&sectionFall2017f)
	db.Create(&sectionFall2017g)
	db.Create(&sectionFall2017h)
	db.Create(&sectionFall2017i)
	db.Create(&sectionFall2017j)
	db.Create(&sectionFall2017k)


	//Spring 2017 sections
	sectionSpring2017a := model.Section{CourseSectionNumber:001, CourseID:course1.CourseID, FacultyID:faculty1.FacultyID,TimeSlotID:timeslotSpring2017.TimeSlotID, LocationID:location4.LocationID}
	sectionSpring2017b := model.Section{CourseSectionNumber:002, CourseID:course2.CourseID, FacultyID:faculty2.FacultyID,TimeSlotID:timeslotSpring12017.TimeSlotID, LocationID:location5.LocationID}
	sectionSpring2017c := model.Section{CourseSectionNumber:001, CourseID:course3.CourseID, FacultyID:faculty3.FacultyID,TimeSlotID:timeslotSpring22017.TimeSlotID, LocationID:location6.LocationID}
	sectionSpring2017d := model.Section{CourseSectionNumber:002, CourseID:course4.CourseID, FacultyID:faculty4.FacultyID,TimeSlotID:timeslotSpring32017.TimeSlotID, LocationID:location7.LocationID}
	sectionSpring2017e := model.Section{CourseSectionNumber:001, CourseID:course5.CourseID, FacultyID:faculty5.FacultyID,TimeSlotID:timeslotSpring42017.TimeSlotID, LocationID:location8.LocationID}
	sectionSpring2017f := model.Section{CourseSectionNumber:003, CourseID:course1.CourseID, FacultyID:faculty6.FacultyID,TimeSlotID:timeslotSpring2017tr.TimeSlotID, LocationID:location9.LocationID}
	sectionSpring2017g := model.Section{CourseSectionNumber:002, CourseID:course2.CourseID, FacultyID:faculty1.FacultyID,TimeSlotID:timeslotSpring2017tr.TimeSlotID, LocationID:location10.LocationID}
	sectionSpring2017h := model.Section{CourseSectionNumber:001, CourseID:course3.CourseID, FacultyID:faculty2.FacultyID,TimeSlotID:timeslotSpring2017tr.TimeSlotID, LocationID:location1.LocationID}
	sectionSpring2017i := model.Section{CourseSectionNumber:002, CourseID:course4.CourseID, FacultyID:faculty3.FacultyID,TimeSlotID:timeslotSpring2017tr.TimeSlotID, LocationID:location2.LocationID}
	sectionSpring2017j := model.Section{CourseSectionNumber:003, CourseID:course5.CourseID, FacultyID:faculty4.FacultyID,TimeSlotID:timeslotSpring2017tr.TimeSlotID, LocationID:location3.LocationID}
	sectionSpring2017k := model.Section{CourseSectionNumber:001, CourseID:course1.CourseID, FacultyID:faculty5.FacultyID,TimeSlotID:timeslotSpring2017tr.TimeSlotID, LocationID:location4.LocationID}

	db.Create(&sectionSpring2017a)
	db.Create(&sectionSpring2017b)
	db.Create(&sectionSpring2017c)
	db.Create(&sectionSpring2017d)
	db.Create(&sectionSpring2017e)
	db.Create(&sectionSpring2017f)
	db.Create(&sectionSpring2017g)
	db.Create(&sectionSpring2017h)
	db.Create(&sectionSpring2017i)
	db.Create(&sectionSpring2017j)
	db.Create(&sectionSpring2017k)


	//enrollment for Spring 2018
	enroll1 := model.Enrollment{StudentID:student1.StudentID, SectionID:section3.SectionID}
	enroll2 := model.Enrollment{StudentID:student2.StudentID, SectionID:section3.SectionID}
	enroll3 := model.Enrollment{StudentID:student3.StudentID, SectionID:section4.SectionID}
	enroll4 := model.Enrollment{StudentID:student4.StudentID, SectionID:section4.SectionID}
	enroll5 := model.Enrollment{StudentID:student1.StudentID, SectionID:section6.SectionID}
	enroll6 := model.Enrollment{StudentID:student2.StudentID, SectionID:section6.SectionID}
	enroll7 := model.Enrollment{StudentID:student3.StudentID, SectionID:section6.SectionID}
	enroll8 := model.Enrollment{StudentID:student4.StudentID, SectionID:section7.SectionID}
	enroll9 := model.Enrollment{StudentID:student1.StudentID, SectionID:section5.SectionID}
	enroll10 := model.Enrollment{StudentID:student2.StudentID, SectionID:section5.SectionID}
	enroll1011 := model.Enrollment{StudentID:student3.StudentID, SectionID:section5.SectionID}
	enroll1012 := model.Enrollment{StudentID:student4.StudentID, SectionID:section5.SectionID}
	enroll11 := model.Enrollment{StudentID:student3.StudentID, SectionID:section1.SectionID}
	enroll12 := model.Enrollment{StudentID:student1.StudentID, SectionID:section2.SectionID}

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
	db.Create(&enroll1011)
	db.Create(&enroll1012)



	//when student registers for a course, a history record must be created as well
	//history for Spring 2018
	history1 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll1.EnrollmentID, Status:"In progress", Grade:"-"}
	history2 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll2.EnrollmentID, Status:"In progress", Grade:"-"}
	history3 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll3.EnrollmentID, Status:"In progress", Grade:"-"}
	history4 := model.StudentHistory{StudentID:student4.StudentID, EnrollmentID:enroll4.EnrollmentID, Status:"In progress", Grade:"-"}
	history5 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll5.EnrollmentID, Status:"In progress", Grade:"-"}
	history6 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll6.EnrollmentID, Status:"In progress", Grade:"-"}
	history7 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll7.EnrollmentID, Status:"In progress", Grade:"-"}
	history8 := model.StudentHistory{StudentID:student4.StudentID, EnrollmentID:enroll8.EnrollmentID, Status:"In progress", Grade:"-"}
	history9 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll9.EnrollmentID, Status:"In progress", Grade:"-"}
	history10 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll10.EnrollmentID, Status:"In progress", Grade:"-"}
	history11 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll11.EnrollmentID, Status:"In progress", Grade:"-"}
	history12 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll12.EnrollmentID, Status:"In progress", Grade:"-"}
	history1222 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll1011.EnrollmentID, Status:"In progress", Grade:"-"}
	history1221 := model.StudentHistory{StudentID:student4.StudentID, EnrollmentID:enroll1012.EnrollmentID, Status:"In progress", Grade:"-"}

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
	db.Create(&history1222)
	db.Create(&history1221)


	//enrollment for Spring 2017
	enroll111 := model.Enrollment{StudentID:student1.StudentID, SectionID:sectionSpring2017a.SectionID}
	enroll21 := model.Enrollment{StudentID:student2.StudentID, SectionID:sectionSpring2017b.SectionID}
	enroll31 := model.Enrollment{StudentID:student3.StudentID, SectionID:sectionSpring2017c.SectionID}
	enroll41 := model.Enrollment{StudentID:student4.StudentID, SectionID:sectionSpring2017d.SectionID}
	enroll51 := model.Enrollment{StudentID:student1.StudentID, SectionID:sectionSpring2017e.SectionID}
	enroll61 := model.Enrollment{StudentID:student2.StudentID, SectionID:sectionSpring2017f.SectionID}
	enroll71 := model.Enrollment{StudentID:student3.StudentID, SectionID:sectionSpring2017g.SectionID}
	enroll81 := model.Enrollment{StudentID:student4.StudentID, SectionID:sectionSpring2017h.SectionID}
	enroll91 := model.Enrollment{StudentID:student1.StudentID, SectionID:sectionSpring2017i.SectionID}
	enroll101 := model.Enrollment{StudentID:student2.StudentID, SectionID:sectionSpring2017j.SectionID}
	enroll1112 := model.Enrollment{StudentID:student3.StudentID, SectionID:sectionSpring2017k.SectionID}
	enroll121 := model.Enrollment{StudentID:student1.StudentID, SectionID:sectionSpring2017b.SectionID}

	db.Create(&enroll111)
	db.Create(&enroll21)
	db.Create(&enroll31)
	db.Create(&enroll41)
	db.Create(&enroll51)
	db.Create(&enroll61)
	db.Create(&enroll71)
	db.Create(&enroll81)
	db.Create(&enroll91)
	db.Create(&enroll101)
	db.Create(&enroll1112)
	db.Create(&enroll121)

	//history Spring 2017
	history11a := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll111.EnrollmentID, Status:"Complete", Grade:"A"}
	history22 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll21.EnrollmentID, Status:"Complete", Grade:"A"}
	history33 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll31.EnrollmentID, Status:"Complete", Grade:"B"}
	history44 := model.StudentHistory{StudentID:student4.StudentID, EnrollmentID:enroll41.EnrollmentID, Status:"Complete", Grade:"B"}
	history55 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll51.EnrollmentID, Status:"Complete", Grade:"B-"}
	history66 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll61.EnrollmentID, Status:"Complete", Grade:"A"}
	history77 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll71.EnrollmentID, Status:"Complete", Grade:"B+"}
	history88 := model.StudentHistory{StudentID:student4.StudentID, EnrollmentID:enroll81.EnrollmentID, Status:"Complete", Grade:"C"}
	history99 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll91.EnrollmentID, Status:"Complete", Grade:"A-"}
	history100 := model.StudentHistory{StudentID:student2.StudentID, EnrollmentID:enroll101.EnrollmentID, Status:"Complete", Grade:"B+"}
	history111 := model.StudentHistory{StudentID:student3.StudentID, EnrollmentID:enroll1112.EnrollmentID, Status:"Dropped", Grade:"-"}
	history122 := model.StudentHistory{StudentID:student1.StudentID, EnrollmentID:enroll121.EnrollmentID, Status:"Complete", Grade:"A"}

	db.Create(&history11a)
	db.Create(&history22)
	db.Create(&history33)
	db.Create(&history44)
	db.Create(&history55)
	db.Create(&history66)
	db.Create(&history77)
	db.Create(&history88)
	db.Create(&history99)
	db.Create(&history100)
	db.Create(&history111)
	db.Create(&history122)

}


