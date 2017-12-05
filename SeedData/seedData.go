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


	//TODO Need seed data for students having majors and minors


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
	db.DropTable(&model.Advises{})
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
		&model.Advises{},

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

	db.Model(&model.Advises{}).AddForeignKey("student_id", "student(student_id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Advises{}).AddForeignKey("faculty_id", "faculty(faculty_id)", "RESTRICT", "RESTRICT")


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
	user11 := model.MainUser{FirstName: "Dwayne", LastName: "Johnson", UserEmail: "Djohn@starfleet.edu", UserPassword: "6yhdf", UserType: 1}
	user12 := model.MainUser{FirstName: "Mike", LastName: "Wayne", UserEmail: "Mwayne@starfleet.edu", UserPassword: "hrer2r", UserType: 1}
	user13 := model.MainUser{FirstName: "Brook", LastName: "Gordon", UserEmail: "Bgordon@starfleet.edu", UserPassword: "cbtdPW", UserType: 1}
	user14 := model.MainUser{FirstName: "Chelsea", LastName: "Hamilton", UserEmail: "Chamilton@starfleet.edu", UserPassword: "oui7h", UserType: 1}
	user15 := model.MainUser{FirstName: "Kacy", LastName: "Yang", UserEmail: "Kyang@starfleet.edu", UserPassword: "74tcn", UserType: 1}
	user16 := model.MainUser{FirstName: "Karen", LastName: "Philips", UserEmail: "Kphil@starfleet.edu", UserPassword: "13rgb", UserType: 1}
	user17 := model.MainUser{FirstName: "Leslie", LastName: "Combs", UserEmail: "Lcombs@starfleet.edu", UserPassword: "24t5y6y", UserType: 1}
	user18 := model.MainUser{FirstName: "Lisa", LastName: "Can", UserEmail: "Lcan@starfleet.edu", UserPassword: "56785y", UserType: 1}
	user19 := model.MainUser{FirstName: "Gabrielle", LastName: "May", UserEmail: "Gmay@starfleet.edu", UserPassword: "566575", UserType: 1}
	user20 := model.MainUser{FirstName: "Nay", LastName: "Books", UserEmail: "Nbooks@starfleet.edu", UserPassword: "jytgfd", UserType: 1}
	user21 := model.MainUser{FirstName: "Kay", LastName: "Bam", UserEmail: "Kbam@starfleet.edu", UserPassword: "434ef", UserType: 1}
	user22 := model.MainUser{FirstName: "Jen", LastName: "Lamb", UserEmail: "Jlamb@starfleet.edu", UserPassword: "bgfht", UserType: 1}
	user23 := model.MainUser{FirstName: "Pam", LastName: "Pen", UserEmail: "Ppen@starfleet.edu", UserPassword: "4htuug", UserType: 1}
	user24 := model.MainUser{FirstName: "Helen", LastName: "Keller", UserEmail: "Hkel@starfleet.edu", UserPassword: "wrg3td", UserType: 1}
	user25 := model.MainUser{FirstName: "Kelly", LastName: "Nevers", UserEmail: "Knevers@starfleet.edu", UserPassword: "h5645s", UserType: 1}
	user26 := model.MainUser{FirstName: "Jill", LastName: "Jack", UserEmail: "Jjack@starfleet.edu", UserPassword: "gwf345h", UserType: 1}
	user27 := model.MainUser{FirstName: "Rell", LastName: "Kell", UserEmail: "Rkell@starfleet.edu", UserPassword: "4h6y54", UserType: 1}
	user28 := model.MainUser{FirstName: "Bell", LastName: "Bank", UserEmail: "Bbank@starfleet.edu", UserPassword: "0swwjv", UserType: 1}
	user29 := model.MainUser{FirstName: "Lola", LastName: "Hank", UserEmail: "Lhank@starfleet.edu", UserPassword: "29whe44", UserType: 1}
	user30 := model.MainUser{FirstName: "Manny", LastName: "Pell", UserEmail: "Mpell@starfleet.edu", UserPassword: "9sgu9g", UserType: 1}
	user31 := model.MainUser{FirstName: "Kenny", LastName: "Ford", UserEmail: "Kford@starfleet.edu", UserPassword: "nkxjnw", UserType: 1}
	user32 := model.MainUser{FirstName: "Nay", LastName: "Books", UserEmail: "Nbooks@starfleet.edu", UserPassword: "jytgfd", UserType: 1}
	user33 := model.MainUser{FirstName: "Frank", LastName: "Brick", UserEmail: "Fmignon@starfleet.edu", UserPassword: "jelly", UserType: 1}
	user34 := model.MainUser{FirstName: "Jonathan", LastName: "Poke", UserEmail: "Bcheese@starfleet.edu", UserPassword: "kwodss", UserType: 1}
	user35 := model.MainUser{FirstName: "Ken", LastName: "Will", UserEmail: "Kwill@starfleet.edu", UserPassword: "svrww", UserType: 1}
	user36 := model.MainUser{FirstName: "Ill", LastName: "Dill", UserEmail: "Illdill@starfleet.edu", UserPassword: "snviuwwwj", UserType: 1}
	user37 := model.MainUser{FirstName: "Kayton", LastName: "Booking", UserEmail: "Kbooking@starfleet.edu", UserPassword: "89jsvnjk", UserType: 1}
	user38 := model.MainUser{FirstName: "Jaleel", LastName: "Neal", UserEmail: "Jneal@starfleet.edu", UserPassword: "kjnwkjnw", UserType: 1}
	user39 := model.MainUser{FirstName: "Wendy", LastName: "Gem", UserEmail: "Wgem@starfleet.edu", UserPassword: "o2ifjwe", UserType: 1}
	user40 := model.MainUser{FirstName: "Sally", LastName: "Mae", UserEmail: "Smae@starfleet.edu", UserPassword: "knswlwe", UserType: 1}
	user41 := model.MainUser{FirstName: "Trey", LastName: "Valley", UserEmail: "Tvalley@starfleet.edu", UserPassword: "sjoi2jw", UserType: 1}
	user42 := model.MainUser{FirstName: "Krystal", LastName: "Clear", UserEmail: "Kclear@starfleet.edu", UserPassword: "jwjknwejk", UserType: 1}
	user43 := model.MainUser{FirstName: "Mona", LastName: "Scott", UserEmail: "Mscott@starfleet.edu", UserPassword: "ers[lw[2", UserType: 1}
	user44 := model.MainUser{FirstName: "Rich", LastName: "Man", UserEmail: "Rmsn@starfleet.edu", UserPassword: "opakve", UserType: 1}
	user45 := model.MainUser{FirstName: "Kay", LastName: "Winn", UserEmail: "Kwinn@starfleet.edu", UserPassword: "slsjios", UserType: 1}
	user46 := model.MainUser{FirstName: "Lens", LastName: "Wilson", UserEmail: "Lwil@starfleet.edu", UserPassword: "niji2o", UserType: 1}
	user47 := model.MainUser{FirstName: "Gong", LastName: "Ho", UserEmail: "Gho@starfleet.edu", UserPassword: "93u94fhiu34", UserType: 1}
	user48 := model.MainUser{FirstName: "Pal", LastName: "Cakes", UserEmail: "Pcakes@starfleet.edu", UserPassword: "9fhewjkwe", UserType: 1}
	user49 := model.MainUser{FirstName: "Pat", LastName: "Megan", UserEmail: "Pmegan@starfleet.edu", UserPassword: "wiuhfui2", UserType: 1}
	user50 := model.MainUser{FirstName: "Nana", LastName: "Yaw", UserEmail: "Nyaw@starfleet.edu", UserPassword: "kjecnjn2", UserType: 1}

	facultyUser1 := model.MainUser{FirstName: "Brandon", LastName: "Sanderson", UserEmail: "sanderson@gmail.com", UserPassword: "nfe435", UserType: 2}
	facultyUser2 := model.MainUser{FirstName: "Mike", LastName: "Tyson", UserEmail: "tyson@gmail.com", UserPassword: "j76755", UserType: 2}
	facultyUser3 := model.MainUser{FirstName: "Eddard", LastName: "Stark", UserEmail: "estark@winterfell.com", UserPassword: "65745", UserType: 2}
	facultyUser4 := model.MainUser{FirstName: "Angela", LastName: "Gordon", UserEmail: "Agordon@winterfell.com", UserPassword: "fr643d", UserType: 2}
	facultyUser5 := model.MainUser{FirstName: "Blue", LastName: "West", UserEmail: "Bwestk@winterfell.com", UserPassword: "jtyt656", UserType: 2}
	facultyUser6 := model.MainUser{FirstName: "Jada", LastName: "Los", UserEmail: "Jlos@winterfell.com", UserPassword: "e45bd56", UserType: 2}
	facultyUser7 := model.MainUser{FirstName: "Prince", LastName: "King", UserEmail: "Pking@winterfell.com", UserPassword: "85894j", UserType: 2}
	facultyUser8 := model.MainUser{FirstName: "Princess", LastName: "June", UserEmail: "Pjune@winterfell.com", UserPassword: "9ehve", UserType: 2}
	facultyUser9 := model.MainUser{FirstName: "Gifty", LastName: "Boateng", UserEmail: "Pboa@winterfell.com", UserPassword: "isjv39", UserType: 2}
	facultyUser10 := model.MainUser{FirstName: "Puff", LastName: "Dad", UserEmail: "Pdad@winterfell.com", UserPassword: "93hjiw", UserType: 2}
	facultyUser11 := model.MainUser{FirstName: "Mase", LastName: "Cam", UserEmail: "Mcam@winterfell.com", UserPassword: "ewweke", UserType: 2}
	facultyUser12 := model.MainUser{FirstName: "Curtis", LastName: "Jackson", UserEmail: "Cjack@winterfell.com", UserPassword: "ofjewk", UserType: 2}
	facultyUser13 := model.MainUser{FirstName: "Jim", LastName: "Slimmy", UserEmail: "Jslim@winterfell.com", UserPassword: "93ifje", UserType: 2}
	facultyUser14 := model.MainUser{FirstName: "Gordon", LastName: "Paul", UserEmail: "Gpaul@winterfell.com", UserPassword: "woj2io2", UserType: 2}
	facultyUser15 := model.MainUser{FirstName: "Missy", LastName: "Elliot", UserEmail: "Melliot@winterfell.com", UserPassword: "odjco22", UserType: 2}
	facultyUser16 := model.MainUser{FirstName: "Cardi", LastName: "Kim", UserEmail: "Ckim@winterfell.com", UserPassword: "oij20fwe", UserType: 2}
	facultyUser17 := model.MainUser{FirstName: "Layla", LastName: "Ali", UserEmail: "Lali@winterfell.com", UserPassword: "ooo2jnfk2", UserType: 2}
	facultyUser18 := model.MainUser{FirstName: "Tyreke", LastName: "Evans", UserEmail: "Tevans@winterfell.com", UserPassword: "owkop2k22", UserType: 2}
	facultyUser19 := model.MainUser{FirstName: "Jenna", LastName: "Dame", UserEmail: "Jdame@winterfell.com", UserPassword: "02i3joe", UserType: 2}
	facultyUser20 := model.MainUser{FirstName: "Eric", LastName: "Iverson", UserEmail: "Eiverson@winterfell.com", UserPassword: "mckmpo2k3", UserType: 2}


	/*
	facultyUser4 := model.MainUser{FirstName: "Brandon", LastName: "Sanderson", UserEmail: "sanderson@gmail.com", UserPassword: "testPW", UserType: 2}
	facultyUser5 := model.MainUser{FirstName: "Mike", LastName: "Tyson", UserEmail: "tyson@gmail.com", UserPassword: "testPW", UserType: 2}
	facultyUser6 := model.MainUser{FirstName: "Eddard", LastName: "Stark", UserEmail: "estark@winterfell.com", UserPassword: "testPW", UserType: 2}
	*/
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
	db.Create(&user11)
	db.Create(&user12)
	db.Create(&user13)
	db.Create(&user14)
	db.Create(&user15)
	db.Create(&user16)
	db.Create(&user17)
	db.Create(&user18)
	db.Create(&user19)
	db.Create(&user20)
	db.Create(&user21)
	db.Create(&user22)
	db.Create(&user23)
	db.Create(&user24)
	db.Create(&user25)
	db.Create(&user26)
	db.Create(&user27)
	db.Create(&user28)
	db.Create(&user29)
	db.Create(&user30)
	db.Create(&user31)
	db.Create(&user32)
	db.Create(&user33)
	db.Create(&user34)
	db.Create(&user35)
	db.Create(&user36)
	db.Create(&user37)
	db.Create(&user38)
	db.Create(&user39)
	db.Create(&user40)
	db.Create(&user41)
	db.Create(&user42)
	db.Create(&user43)
	db.Create(&user44)
	db.Create(&user45)
	db.Create(&user46)
	db.Create(&user47)
	db.Create(&user48)
	db.Create(&user49)
	db.Create(&user50)

	db.Create(&facultyUser1)
	db.Create(&facultyUser2)
	db.Create(&facultyUser3)
	db.Create(&facultyUser4)
	db.Create(&facultyUser5)
	db.Create(&facultyUser6)
	db.Create(&facultyUser7)
	db.Create(&facultyUser8)
	db.Create(&facultyUser9)
	db.Create(&facultyUser10)
	db.Create(&facultyUser11)
	db.Create(&facultyUser12)
	db.Create(&facultyUser13)
	db.Create(&facultyUser14)
	db.Create(&facultyUser15)
	db.Create(&facultyUser16)
	db.Create(&facultyUser17)
	db.Create(&facultyUser18)
	db.Create(&facultyUser19)
	db.Create(&facultyUser20)



	student2 := model.Student{StudentID: user2.UserID, StudentType: 1}
	student1 := model.Student{StudentID: user1.UserID, StudentType: 1}
	student3 := model.Student{StudentID: user3.UserID, StudentType: 2}
	student4 := model.Student{StudentID: user4.UserID, StudentType: 2}
	student10 := model.Student{StudentID: user10.UserID, StudentType: 1}
	student11 := model.Student{StudentID: user11.UserID, StudentType: 2}
	student12 := model.Student{StudentID: user12.UserID, StudentType: 2}
	student13 := model.Student{StudentID: user13.UserID, StudentType: 2}
	student14 := model.Student{StudentID: user14.UserID, StudentType: 2}
	student15 := model.Student{StudentID: user15.UserID, StudentType: 2}
	student16 := model.Student{StudentID: user16.UserID, StudentType: 2}
	student17 := model.Student{StudentID: user17.UserID, StudentType: 2}
	student18 := model.Student{StudentID: user18.UserID, StudentType: 2}
	student19 := model.Student{StudentID: user19.UserID, StudentType: 2}
	student20 := model.Student{StudentID: user20.UserID, StudentType: 2}
	student21 := model.Student{StudentID: user21.UserID, StudentType: 1}
	student22 := model.Student{StudentID: user22.UserID, StudentType: 1}
	student23 := model.Student{StudentID: user23.UserID, StudentType: 1}
	student24 := model.Student{StudentID: user24.UserID, StudentType: 1}
	student25 := model.Student{StudentID: user25.UserID, StudentType: 1}
	student26 := model.Student{StudentID: user26.UserID, StudentType: 1}
	student27 := model.Student{StudentID: user27.UserID, StudentType: 1}
	student28 := model.Student{StudentID: user28.UserID, StudentType: 1}
	student29 := model.Student{StudentID: user29.UserID, StudentType: 1}
	student30 := model.Student{StudentID: user30.UserID, StudentType: 1}
	student31 := model.Student{StudentID: user31.UserID, StudentType: 1}
	student32 := model.Student{StudentID: user32.UserID, StudentType: 1}
	student33 := model.Student{StudentID: user33.UserID, StudentType: 1}
	student34 := model.Student{StudentID: user34.UserID, StudentType: 1}
	student35 := model.Student{StudentID: user35.UserID, StudentType: 1}
	student36 := model.Student{StudentID: user36.UserID, StudentType: 1}
	student37 := model.Student{StudentID: user37.UserID, StudentType: 1}
	student38 := model.Student{StudentID: user38.UserID, StudentType: 1}
	student39 := model.Student{StudentID: user39.UserID, StudentType: 1}
	student40 := model.Student{StudentID: user40.UserID, StudentType: 1}
	student41 := model.Student{StudentID: user41.UserID, StudentType: 1}
	student42 := model.Student{StudentID: user42.UserID, StudentType: 1}
	student43 := model.Student{StudentID: user43.UserID, StudentType: 1}
	student44 := model.Student{StudentID: user44.UserID, StudentType: 1}
	student45 := model.Student{StudentID: user45.UserID, StudentType: 1}
	student46 := model.Student{StudentID: user46.UserID, StudentType: 1}
	student47 := model.Student{StudentID: user47.UserID, StudentType: 1}
	student48 := model.Student{StudentID: user48.UserID, StudentType: 1}
	student49 := model.Student{StudentID: user49.UserID, StudentType: 1}
	student50 := model.Student{StudentID: user50.UserID, StudentType: 1}

	db.Create(&student1)
	db.Create(&student2)
	db.Create(&student3)
	db.Create(&student4)

	db.Create(&student10)
	db.Create(&student11)
	db.Create(&student12)
	db.Create(&student13)
	db.Create(&student14)
	db.Create(&student15)
	db.Create(&student16)
	db.Create(&student17)
	db.Create(&student18)
	db.Create(&student19)
	db.Create(&student20)
	db.Create(&student21)
	db.Create(&student22)
	db.Create(&student23)
	db.Create(&student24)
	db.Create(&student25)
	db.Create(&student26)
	db.Create(&student27)
	db.Create(&student28)
	db.Create(&student29)
	db.Create(&student30)
	db.Create(&student31)
	db.Create(&student32)
	db.Create(&student33)
	db.Create(&student34)
	db.Create(&student35)
	db.Create(&student36)
	db.Create(&student37)
	db.Create(&student38)
	db.Create(&student39)
	db.Create(&student40)
	db.Create(&student41)
	db.Create(&student42)
	db.Create(&student43)
	db.Create(&student44)
	db.Create(&student45)
	db.Create(&student46)
	db.Create(&student47)
	db.Create(&student48)
	db.Create(&student49)
	db.Create(&student50)


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

	//fullTimeStudent5 := model.FullTimeStudent{FullTimeStudentID: student5.StudentID, NumCredits: 16}
	//fullTimeStudent6 := model.FullTimeStudent{FullTimeStudentID: student6.StudentID, NumCredits: 18}
	//fullTimeStudent7 := model.FullTimeStudent{FullTimeStudentID: student7.StudentID, NumCredits: 16}
	//fullTimeStudent8 := model.FullTimeStudent{FullTimeStudentID: student8.StudentID, NumCredits: 16}
	//fullTimeStudent9 := model.FullTimeStudent{FullTimeStudentID: student9.StudentID, NumCredits: 18}
	fullTimeStudent10 := model.FullTimeStudent{FullTimeStudentID: student10.StudentID, NumCredits: 16}
	db.Create(&fullTimeStudent1)
	db.Create(&fullTimeStudent2)
	//db.Create(&fullTimeStudent3)
	//db.Create(&fullTimeStudent4)
	//db.Create(&fullTimeStudent5)
	//db.Create(&fullTimeStudent6)
	//db.Create(&fullTimeStudent7)
	//db.Create(&fullTimeStudent8)
	//db.Create(&fullTimeStudent9)
	db.Create(&fullTimeStudent10)
	partTimeStudent11 := model.PartTimeStudent{PartTimeStudentID: student11.StudentID, NumCredits: 8}
	partTimeStudent12 := model.PartTimeStudent{PartTimeStudentID: student12.StudentID, NumCredits: 8}
	partTimeStudent13 := model.PartTimeStudent{PartTimeStudentID: student13.StudentID, NumCredits: 4}
	partTimeStudent14 := model.PartTimeStudent{PartTimeStudentID: student14.StudentID, NumCredits: 4}
	partTimeStudent15 := model.PartTimeStudent{PartTimeStudentID: student15.StudentID, NumCredits: 8}
	partTimeStudent16 := model.PartTimeStudent{PartTimeStudentID: student16.StudentID, NumCredits: 8}
	partTimeStudent17 := model.PartTimeStudent{PartTimeStudentID: student17.StudentID, NumCredits: 8}
	partTimeStudent18 := model.PartTimeStudent{PartTimeStudentID: student18.StudentID, NumCredits: 8}
	partTimeStudent19 := model.PartTimeStudent{PartTimeStudentID: student19.StudentID, NumCredits: 4}
	partTimeStudent20 := model.PartTimeStudent{PartTimeStudentID: student20.StudentID, NumCredits: 8}
	db.Create(&partTimeStudent11)
	db.Create(&partTimeStudent12)
	db.Create(&partTimeStudent13)
	db.Create(&partTimeStudent14)
	db.Create(&partTimeStudent15)
	db.Create(&partTimeStudent16)
	db.Create(&partTimeStudent17)
	db.Create(&partTimeStudent18)
	db.Create(&partTimeStudent19)
	db.Create(&partTimeStudent20)
	fullTimeStudent21 := model.FullTimeStudent{FullTimeStudentID: student21.StudentID, NumCredits: 16}
	fullTimeStudent22 := model.FullTimeStudent{FullTimeStudentID: student22.StudentID, NumCredits: 16}
	fullTimeStudent23 := model.FullTimeStudent{FullTimeStudentID: student23.StudentID, NumCredits: 16}
	fullTimeStudent24 := model.FullTimeStudent{FullTimeStudentID: student24.StudentID, NumCredits: 16}
	fullTimeStudent25 := model.FullTimeStudent{FullTimeStudentID: student25.StudentID, NumCredits: 16}
	fullTimeStudent26 := model.FullTimeStudent{FullTimeStudentID: student26.StudentID, NumCredits: 16}
	fullTimeStudent27 := model.FullTimeStudent{FullTimeStudentID: student27.StudentID, NumCredits: 16}
	fullTimeStudent28 := model.FullTimeStudent{FullTimeStudentID: student28.StudentID, NumCredits: 16}
	fullTimeStudent29 := model.FullTimeStudent{FullTimeStudentID: student29.StudentID, NumCredits: 16}
	fullTimeStudent30 := model.FullTimeStudent{FullTimeStudentID: student30.StudentID, NumCredits: 16}
	fullTimeStudent31 := model.FullTimeStudent{FullTimeStudentID: student31.StudentID, NumCredits: 16}
	fullTimeStudent32 := model.FullTimeStudent{FullTimeStudentID: student32.StudentID, NumCredits: 16}
	fullTimeStudent33 := model.FullTimeStudent{FullTimeStudentID: student33.StudentID, NumCredits: 16}
	fullTimeStudent34 := model.FullTimeStudent{FullTimeStudentID: student34.StudentID, NumCredits: 16}
	fullTimeStudent35 := model.FullTimeStudent{FullTimeStudentID: student35.StudentID, NumCredits: 16}
	fullTimeStudent36 := model.FullTimeStudent{FullTimeStudentID: student36.StudentID, NumCredits: 16}
	fullTimeStudent37 := model.FullTimeStudent{FullTimeStudentID: student37.StudentID, NumCredits: 16}
	fullTimeStudent38 := model.FullTimeStudent{FullTimeStudentID: student38.StudentID, NumCredits: 16}
	fullTimeStudent39 := model.FullTimeStudent{FullTimeStudentID: student39.StudentID, NumCredits: 16}
	fullTimeStudent40 := model.FullTimeStudent{FullTimeStudentID: student40.StudentID, NumCredits: 16}
	fullTimeStudent41 := model.FullTimeStudent{FullTimeStudentID: student41.StudentID, NumCredits: 16}
	fullTimeStudent42 := model.FullTimeStudent{FullTimeStudentID: student42.StudentID, NumCredits: 16}
	fullTimeStudent43 := model.FullTimeStudent{FullTimeStudentID: student43.StudentID, NumCredits: 16}
	fullTimeStudent44 := model.FullTimeStudent{FullTimeStudentID: student44.StudentID, NumCredits: 16}
	fullTimeStudent45 := model.FullTimeStudent{FullTimeStudentID: student45.StudentID, NumCredits: 16}
	fullTimeStudent46 := model.FullTimeStudent{FullTimeStudentID: student46.StudentID, NumCredits: 16}
	fullTimeStudent47 := model.FullTimeStudent{FullTimeStudentID: student47.StudentID, NumCredits: 16}
	fullTimeStudent48 := model.FullTimeStudent{FullTimeStudentID: student48.StudentID, NumCredits: 16}
	fullTimeStudent49 := model.FullTimeStudent{FullTimeStudentID: student49.StudentID, NumCredits: 16}
	fullTimeStudent50 := model.FullTimeStudent{FullTimeStudentID: student50.StudentID, NumCredits: 16}
	db.Create(&fullTimeStudent21)
	db.Create(&fullTimeStudent22)
	db.Create(&fullTimeStudent23)
	db.Create(&fullTimeStudent24)
	db.Create(&fullTimeStudent25)
	db.Create(&fullTimeStudent26)
	db.Create(&fullTimeStudent27)
	db.Create(&fullTimeStudent28)
	db.Create(&fullTimeStudent29)
	db.Create(&fullTimeStudent30)
	db.Create(&fullTimeStudent31)
	db.Create(&fullTimeStudent32)
	db.Create(&fullTimeStudent33)
	db.Create(&fullTimeStudent34)
	db.Create(&fullTimeStudent35)
	db.Create(&fullTimeStudent36)
	db.Create(&fullTimeStudent37)
	db.Create(&fullTimeStudent38)
	db.Create(&fullTimeStudent39)
	db.Create(&fullTimeStudent40)
	db.Create(&fullTimeStudent41)
	db.Create(&fullTimeStudent42)
	db.Create(&fullTimeStudent43)
	db.Create(&fullTimeStudent44)
	db.Create(&fullTimeStudent45)
	db.Create(&fullTimeStudent46)
	db.Create(&fullTimeStudent47)
	db.Create(&fullTimeStudent48)
	db.Create(&fullTimeStudent49)
	db.Create(&fullTimeStudent50)

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
	department6 := model.Department{DepartmentName: "Biological Sciences", DepartmentBuilding: "BiologyBuilding", DepartmentRoomNumber: "500", DepartmentChair: "Franklin, Kyle", DepartmentPhoneNumber: "516-543-3591"}
	department7 := model.Department{DepartmentName: "Chemistry", DepartmentBuilding: "ChemistryBuilding", DepartmentRoomNumber: "600", DepartmentChair: "Bennet, Daniel", DepartmentPhoneNumber: "516-633-1121"}
	department8 := model.Department{DepartmentName: "Physics", DepartmentBuilding: "PhysicsBuilding", DepartmentRoomNumber: "700", DepartmentChair: "Bryant, Russell", DepartmentPhoneNumber: "516-223-3331"}
	department9 := model.Department{DepartmentName: "Criminology", DepartmentBuilding: "CriminologyBuilding", DepartmentRoomNumber: "800", DepartmentChair: "George, Dwyane", DepartmentPhoneNumber: "516-993-3451"}
	department10 := model.Department{DepartmentName: "English", DepartmentBuilding: "EnglishBuilding", DepartmentRoomNumber: "450", DepartmentChair: "Wade, Paul", DepartmentPhoneNumber: "516-785-3943"}
	department11 := model.Department{DepartmentName: "Health", DepartmentBuilding: "HealthBuilding", DepartmentRoomNumber: "150", DepartmentChair: "James, Kyrie", DepartmentPhoneNumber: "516-330-3491"}
	department12 := model.Department{DepartmentName: "Music", DepartmentBuilding: "MusicBuilding", DepartmentRoomNumber: "250", DepartmentChair: "Davis, Earl", DepartmentPhoneNumber: "516-432-1351"}
	department13 := model.Department{DepartmentName: "Psychology", DepartmentBuilding: "PsychologyBuilding", DepartmentRoomNumber: "350", DepartmentChair: "Curry, Kevin", DepartmentPhoneNumber: "516-995-3491"}
	department14 := model.Department{DepartmentName: "Sociology", DepartmentBuilding: "SociologyBuilding", DepartmentRoomNumber: "550", DepartmentChair: "Thompson, John", DepartmentPhoneNumber: "516-785-6977"}
	department15 := model.Department{DepartmentName: "Education", DepartmentBuilding: "EducationBuilding", DepartmentRoomNumber: "650", DepartmentChair: "Butler, Klay", DepartmentPhoneNumber: "516-069-0000"}


	fmt.Println("Creating departments")

	db.Create(&department1)
	db.Create(&department2)
	db.Create(&department3)
	db.Create(&department4)
	db.Create(&department5)
	db.Create(&department6)
	db.Create(&department7)
	db.Create(&department8)
	db.Create(&department9)
	db.Create(&department10)
	db.Create(&department11)
	db.Create(&department12)
	db.Create(&department13)
	db.Create(&department14)
	db.Create(&department15)

	faculty1 := model.Faculty{FacultyID: user5.UserID, FacultyType: 1, RoomNumber:"B100", DepartmentID: department1.DepartmentID}
	faculty2 := model.Faculty{FacultyID: user6.UserID, FacultyType: 2, RoomNumber:"C300",DepartmentID: department2.DepartmentID}
	faculty3 := model.Faculty{FacultyID: user7.UserID, FacultyType: 1, RoomNumber:"B110",DepartmentID: department1.DepartmentID}

	faculty4 := model.Faculty{FacultyID: facultyUser4.UserID, FacultyType: 1, RoomNumber:"D100", DepartmentID: department3.DepartmentID}
	faculty5 := model.Faculty{FacultyID: facultyUser5.UserID, FacultyType: 1, RoomNumber:"B200", DepartmentID: department4.DepartmentID}
	faculty6 := model.Faculty{FacultyID: facultyUser6.UserID, FacultyType: 1, RoomNumber:"B400", DepartmentID: department5.DepartmentID}


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


	advisor1 := model.Advises{FacultyID:faculty1.FacultyID,StudentID:student1.StudentID}
	advisor2 := model.Advises{FacultyID:faculty1.FacultyID,StudentID:student2.StudentID}
	advisor3 := model.Advises{FacultyID:faculty3.FacultyID,StudentID:student3.StudentID}
	advisor4 := model.Advises{FacultyID:faculty4.FacultyID,StudentID:student4.StudentID}

	db.Create(&advisor1)
	db.Create(&advisor2)
	db.Create(&advisor3)
	db.Create(&advisor4)

	major1 := model.Major{DepartmentID: department1.DepartmentID, MajorName: "Math"}
	major2 := model.Major{DepartmentID: department1.DepartmentID, MajorName: "Physics"}
	major3 := model.Major{DepartmentID: department6.DepartmentID, MajorName: "Biological Sciences"}
	major4 := model.Major{DepartmentID: department10.DepartmentID, MajorName: "English"}
	major5 := model.Major{DepartmentID: department3.DepartmentID, MajorName: "History"}
	major6 := model.Major{DepartmentID: department14.DepartmentID, MajorName: "Sociology"}
	major7 := model.Major{DepartmentID: department13.DepartmentID, MajorName: "Psychology"}
	major8 := model.Major{DepartmentID: department2.DepartmentID, MajorName: "Accounting"}
	major9 := model.Major{DepartmentID: department7.DepartmentID, MajorName: "Chemistry"}
	major10 := model.Major{DepartmentID: department2.DepartmentID, MajorName: "Computer Science"}
	major11 := model.Major{DepartmentID: department9.DepartmentID, MajorName: "Criminology"}
	major12 := model.Major{DepartmentID: department15.DepartmentID, MajorName: "Philosophy"}
	major13 := model.Major{DepartmentID: department4.DepartmentID, MajorName: "Spanish"}
	major14 := model.Major{DepartmentID: department4.DepartmentID, MajorName: "French"}
	major15 := model.Major{DepartmentID: department4.DepartmentID, MajorName: "liberal Arts"}


	db.Create(&major1)
	db.Create(&major2)
	db.Create(&major3)
	db.Create(&major4)
	db.Create(&major5)
	db.Create(&major6)
	db.Create(&major7)
	db.Create(&major8)
	db.Create(&major9)
	db.Create(&major10)
	db.Create(&major11)
	db.Create(&major12)
	db.Create(&major13)
	db.Create(&major14)
	db.Create(&major15)

	minor1 := model.Minor{DepartmentID: department1.DepartmentID, MinorName: "APPLIED MATH"}
	minor2 := model.Minor{DepartmentID: department1.DepartmentID, MinorName: "MINOR MATH"}
	minor3 := model.Minor{DepartmentID: department2.DepartmentID, MinorName: "LITTLE COMPUTERS"}
	minor4 := model.Minor{DepartmentID: department2.DepartmentID, MinorName: "COMPUTER STUFF"}
	minor5 := model.Minor{DepartmentID: department2.DepartmentID, MinorName: "Finance"}
	db.Create(&minor1)
	db.Create(&minor2)
	db.Create(&minor3)
	db.Create(&minor4)
	db.Create(&minor5)

	hold1 := model.Hold{HoldName: "Unpaid Bill"}
	hold2 := model.Hold{HoldName: "Un-submitted Health Forms"}
	hold3 := model.Hold{HoldName: "Unpaid Parking Ticket"}
	hold4 := model.Hold{HoldName: "Unpaid Room Charges"}
	hold5 := model.Hold{HoldName: "Unpaid Tuition"}
	hold6 := model.Hold{HoldName: "Library Fines"}
	hold7 := model.Hold{HoldName: "Unpaid Meal Plans"}
	hold8 := model.Hold{HoldName: "Non-return of athletic Equipment"}
	hold9 := model.Hold{HoldName: "Un-submitted Financial Aid Forms"}
	hold10 := model.Hold{HoldName: "Registration"}

	db.Create(&hold1)
	db.Create(&hold2)
	db.Create(&hold3)
	db.Create(&hold4)
	db.Create(&hold5)
	db.Create(&hold6)
	db.Create(&hold7)
	db.Create(&hold8)
	db.Create(&hold9)
	db.Create(&hold10)

	studenthold1 := model.StudentHolds{StudentID: student1.StudentID, HoldID: hold1.HoldID}
	studenthold2 := model.StudentHolds{StudentID: student1.StudentID, HoldID: hold2.HoldID}
	studenthold3 := model.StudentHolds{StudentID: student2.StudentID, HoldID: hold3.HoldID}
	studenthold4 := model.StudentHolds{StudentID: student1.StudentID, HoldID: hold4.HoldID}

	db.Create(&studenthold1)
	db.Create(&studenthold2)
	db.Create(&studenthold3)
	db.Create(&studenthold4)

	course1 := model.Course{CourseName: "Warp Field Mechanics", CourseCredits: 4, DepartmentID: department1.DepartmentID, CourseDescription: "An introduction to the theory behind faster than light space travel."}
	course2 := model.Course{CourseName: "History of Space Travel", CourseCredits: 4, DepartmentID: department3.DepartmentID, CourseDescription: "A survey of humanity's entry into the age of space exploration, from the first artificial satellite to first contact with the Vulcans. "}
	course3 := model.Course{CourseName: "Contemporary Holography", CourseCredits: 4, DepartmentID: department4.DepartmentID, CourseDescription: "An introduction to Holography."}
	course4 := model.Course{CourseName: "Chemistry I", CourseCredits: 4, DepartmentID: department7.DepartmentID, CourseDescription: "Basic principles to be covered include atomic structure, chemical properties, chemical reactions, the nature of organic and inorganic compounds and novel materials."}
	course5 := model.Course{CourseName: "Accounting I", CourseCredits: 4, DepartmentID: department5.DepartmentID, CourseDescription: "Inroduction to financial accounting with focus on how financial statements are structured and used by proprietorships and cooperations."}
	course6 := model.Course{CourseName: "Microeconomics", CourseCredits: 4, DepartmentID: department5.DepartmentID, CourseDescription: "An introduction to conventional macroeconomic theory, including the determination of national income, government taxing and spending policy, money and banking, unemployment, inflation, economic growth and international trade."}
	course7 := model.Course{CourseName: "Programming I", CourseCredits: 4, DepartmentID: department2.DepartmentID, CourseDescription: "Introduction to program design and analysis: algorithmic processes, basic programming techniques, program specification & structure, program development, debugging, and testing."}
	course8 := model.Course{CourseName: "Calculus I", CourseCredits: 4, DepartmentID: department1.DepartmentID, CourseDescription: "Limits, differentiation, and integration. Relevant applications from the areas of business, economics and the social sciences."}
	course9 := model.Course{CourseName: "French I", CourseCredits: 4, DepartmentID: department4.DepartmentID, CourseDescription: "Each provides for acquisition of the communication skills in French, with special emphasis on aural comprehension and speaking."}
	course10 := model.Course{CourseName: "Spanish I", CourseCredits: 4, DepartmentID: department4.DepartmentID, CourseDescription: "Communication skills in standard Spanish are stressed with particular emphasis placed on speaking and aural comprehension."}
	course11 := model.Course{CourseName: "Intro to Psychology", CourseCredits: 4, DepartmentID: department13.DepartmentID, CourseDescription: "Based on methods and data of scientific psychology, this course investigates basic principles of behavior."}
	course12 := model.Course{CourseName: "Intro to Sociology", CourseCredits: 4, DepartmentID: department14.DepartmentID, CourseDescription: "Such major social institutions as the family, education, politics, law, media, and religion are examined, together with such social processes as socialization, social change, social control, and social stratification."}
	course13 := model.Course{CourseName: "Language Arts", CourseCredits: 4, DepartmentID: department10.DepartmentID, CourseDescription: "The purpose of this course is to explore techniques and strategies to integrate reading instructions across the content areas, with special emphasis on Social Studies."}
	course14 := model.Course{CourseName: "Biology I", CourseCredits: 4, DepartmentID: department6.DepartmentID, CourseDescription: "Surveys the major concepts and principles of biology, including cell structure and function, genetics, ecology, diversity and evolution."}
	course15 := model.Course{CourseName: "Physical Science", CourseCredits: 4, DepartmentID: department8.DepartmentID, CourseDescription: "Fundamental principles of physics, astronomy, chemstry & earth science are covered to provide the student with a broad basic background."}
	course16 := model.Course{CourseName: "Intro to Criminology", CourseCredits: 4, DepartmentID: department9.DepartmentID, CourseDescription: "This course will provide students with an introduction to the social scientific study of crime."}
	course17 := model.Course{CourseName: "Computer Programming II", CourseCredits: 4, DepartmentID: department2.DepartmentID, CourseDescription: "Discussion of storage classes, pointers, recursion, files and string manipulation."}
	course18 := model.Course{CourseName: "World History I: Non-Western", CourseCredits: 4, DepartmentID: department3.DepartmentID, CourseDescription: "This course surveys the history and culture of five regions of the world: the Middle East, sub-Saharan Africa, China, India, and Japan."}
	course19 := model.Course{CourseName: "Biology II", CourseCredits: 4, DepartmentID: department6.DepartmentID, CourseDescription: "The topics considered include: introduction to the chemistry of life, cell biology, genetics, evolution, ecology, cellular and organismic cellular and organismic reproduction, comparative anatomy and the physiology of selected organ systems with emphasis on humans."}
	course20 := model.Course{CourseName: "Chemistry II ", CourseCredits: 4, DepartmentID: department7.DepartmentID, CourseDescription: "Topics include reactions of hydrocarbons, alcohols,ethers, amines and carbonyl compounds, structure and metobolism of carbohydrates, lipids, proteins, enzymes and important body fluids."}
	course21 := model.Course{CourseName: "General Physics I", CourseCredits: 4, DepartmentID: department8.DepartmentID, CourseDescription: "The first semester of a calculus-based introductory physics course primary for chemistry, math, and pre-engineering majors."}
	course22 := model.Course{CourseName: "Drugs and Society", CourseCredits: 4, DepartmentID: department9.DepartmentID, CourseDescription: "This course examines the use of drugs not only in contemporary American society, but also globally, and satisfies the Sociology Department's cross-cultural component."}
	course23 := model.Course{CourseName: "Advanced Composition", CourseCredits: 4, DepartmentID: department10.DepartmentID, CourseDescription: "This course will explore the psychology of language and the relations among languages, behavior, and cognitive processes."}
	course24 := model.Course{CourseName: "Intro to Health & Society", CourseCredits: 4, DepartmentID: department11.DepartmentID, CourseDescription: "Introduces students to the public health perspective on health, also called the social determinants of health, which includes a detailed examination of social class, racism, gender, community, environment."}
	course25 := model.Course{CourseName: "Women & Health", CourseCredits: 4, DepartmentID: department11.DepartmentID, CourseDescription: "TA critical look at women's health issues from the pesrpective of the womens health movement."}
	course26 := model.Course{CourseName: "Music of Global Cultures I", CourseCredits: 4, DepartmentID: department12.DepartmentID, CourseDescription: "An introductory course focused on understanding and exploring the meaning and art of traditional and popular music and global cultures."}
	course27 := model.Course{CourseName: "Development of American Jazz", CourseCredits: 4, DepartmentID: department12.DepartmentID, CourseDescription: "It is a true manifestation of American ideals: democracy in action, freedom of choice through improvisation, and a national identity in music for African Americans."}
	course28 := model.Course{CourseName: "Children's Literature", CourseCredits: 4, DepartmentID: department15.DepartmentID, CourseDescription: "An introduction to the study of literature for children in the pre-school and elementary grades."}
	course29 := model.Course{CourseName: "Issues in Multi-Cultural Ed", CourseCredits: 4, DepartmentID: department15.DepartmentID, CourseDescription: "Focuses on the educational needs of children of varying ethnic, cultural, and/or language backgrounds."}
	course30 := model.Course{CourseName: "Intro to Criminology", CourseCredits: 4, DepartmentID: department2.DepartmentID, CourseDescription: "This course will provide students with an introduction to the social scientific study of crime."}
	course31 := model.Course{CourseName: "Systems Design & Implement", CourseCredits: 4, DepartmentID: department2.DepartmentID, CourseDescription: "System feasibilty studies, meeting with users, project definition, environmental analysis."}
	course32 := model.Course{CourseName: "Probability and Statistics", CourseCredits: 4, DepartmentID: department1.DepartmentID, CourseDescription: "This course presents the m athematical laws of random phenomena, including discrete and continuous random variables, expectation and variance, and common probability distributions such as the binomial, Poisson, and normal distributions."}
	course33 := model.Course{CourseName: "Quantum Chemistry", CourseCredits: 4, DepartmentID: department7.DepartmentID, CourseDescription: "An introduction to the mathematical and physical principles of quatum chemistry, including vector spaces, operator algebra, matrix theory"}
	course34 := model.Course{CourseName: "Senior Seminar in Criminology", CourseCredits: 4, DepartmentID: department9.DepartmentID, CourseDescription: "Students will read and review a variety of major works in contemporary Criminology."}
	course35 := model.Course{CourseName: "Senior Seminar I: Methodology", CourseCredits: 4, DepartmentID: department10.DepartmentID, CourseDescription: "Autobiographical and biographical readings in the intellectual development aid students in understanding issues of personal and career identity, the impact of literary study on the formation of self, and related topics."}
	course36 := model.Course{CourseName: "Environmental Policy&Politics", CourseCredits: 4, DepartmentID: department11.DepartmentID, CourseDescription: "Focuses on the most important federal, environmental, state and local laws."}
	course37 := model.Course{CourseName: "Pre-Student Teaching Seminar", CourseCredits: 4, DepartmentID: department15.DepartmentID, CourseDescription: "This course is designed for Middle Childhood (5-9) & Adolescence Education (7-12) students to comply with state mandated pre-student teaching requirements."}
	course38 := model.Course{CourseName: "Issues in Sociology", CourseCredits: 4, DepartmentID: department14.DepartmentID, CourseDescription: "Gives students the opportunity to examine, in greater depth, the analysis and discussion of current specialized sociological work in the particular area defined by the instructor."}
	course39 := model.Course{CourseName: "Industrial Sociology", CourseCredits: 4, DepartmentID: department14.DepartmentID, CourseDescription: "Examines, from a sociological perspective, the meaning and functions of work in the United States."}
	course40 := model.Course{CourseName: "Political Economy of Africa", CourseCredits: 4, DepartmentID: department5.DepartmentID, CourseDescription: "Explores patterns of politics and issues related to political power and social change in contemporary Africa."}


	db.Create(&course1)
	db.Create(&course2)
	db.Create(&course3)
	db.Create(&course4)
	db.Create(&course5)
	db.Create(&course6)
	db.Create(&course7)
	db.Create(&course8)
	db.Create(&course9)
	db.Create(&course10)
	db.Create(&course11)
	db.Create(&course12)
	db.Create(&course13)
	db.Create(&course14)
	db.Create(&course15)
	db.Create(&course16)
	db.Create(&course17)
	db.Create(&course18)
	db.Create(&course19)
	db.Create(&course20)
	db.Create(&course21)
	db.Create(&course22)
	db.Create(&course23)
	db.Create(&course24)
	db.Create(&course25)
	db.Create(&course26)
	db.Create(&course27)
	db.Create(&course28)
	db.Create(&course29)
	db.Create(&course30)
	db.Create(&course31)
	db.Create(&course32)
	db.Create(&course33)
	db.Create(&course34)
	db.Create(&course35)
	db.Create(&course36)
	db.Create(&course37)
	db.Create(&course38)
	db.Create(&course39)
	db.Create(&course40)

	preReq1 := model.Prerequisite{CourseRequiredBy: course5.CourseID, CourseRequirement: course4.CourseID}
	preReq2 := model.Prerequisite{CourseRequiredBy: course1.CourseID, CourseRequirement: course2.CourseID}
	db.Create(&preReq1)
	db.Create(&preReq2)

	building := model.Building{BuildingName: "The Academy", BuildingAddress: "5 Shawsington Road"}
	building2 := model.Building{BuildingName: "Riften Building", BuildingAddress: "115 Shawsington Road"}
	building3 := model.Building{BuildingName: "Star Building", BuildingAddress: "100 Shawsington Road"}

	db.Create(&building)
	db.Create(&building2)
	db.Create(&building3)


	//building 1
	room1 := model.Room{RoomNumber: "B100", RoomType: "Lecture Hall", RoomCapacity: 100}
	room2 := model.Room{RoomNumber: "C200", RoomType: "LAB", RoomCapacity: 10}
	room3 := model.Room{RoomNumber: "C210", RoomType: "Classroom", RoomCapacity: 30}
	room4 := model.Room{RoomNumber: "C220", RoomType: "Classroom", RoomCapacity: 30}
	room5 := model.Room{RoomNumber: "C230", RoomType: "Classroom", RoomCapacity: 40}
	room6 := model.Room{RoomNumber: "C240", RoomType: "Classroom", RoomCapacity: 20}
	room7 := model.Room{RoomNumber: "C250", RoomType: "Classroom", RoomCapacity: 35}
	room8 := model.Room{RoomNumber: "C260", RoomType: "Classroom", RoomCapacity: 25}
	room9 := model.Room{RoomNumber: "C270", RoomType: "Classroom", RoomCapacity: 25}
	room10 := model.Room{RoomNumber: "C280", RoomType: "Classroom", RoomCapacity: 40}

	//building 2
	b2room1 := model.Room{RoomNumber: "P100", RoomType: "Classroom", RoomCapacity: 25}
	b2room2 := model.Room{RoomNumber: "P150", RoomType: "Classroom", RoomCapacity: 25}
	b2room3 := model.Room{RoomNumber: "P180", RoomType: "Classroom", RoomCapacity: 30}
	b2room4 := model.Room{RoomNumber: "H100", RoomType: "Lecture Hall", RoomCapacity: 125}
	b2room5 := model.Room{RoomNumber: "P180", RoomType: "Classroom", RoomCapacity: 40}
	b2room6 := model.Room{RoomNumber: "P180", RoomType: "Classroom", RoomCapacity: 40}
	b2room7 := model.Room{RoomNumber: "P180", RoomType: "Classroom", RoomCapacity: 20}
	b2room8 := model.Room{RoomNumber: "P180", RoomType: "Classroom", RoomCapacity: 20}
	b2room9 := model.Room{RoomNumber: "P180", RoomType: "Classroom", RoomCapacity: 20}
	b2room10 := model.Room{RoomNumber: "H190", RoomType: "Lecture Hall", RoomCapacity: 125}

	//building 3
	b3room1 := model.Room{RoomNumber: "D100", RoomType: "Classroom", RoomCapacity: 15}
	b3room2 := model.Room{RoomNumber: "D150", RoomType: "Classroom", RoomCapacity: 25}
	b3room3 := model.Room{RoomNumber: "D380", RoomType: "Classroom", RoomCapacity: 10}
	b3room4 := model.Room{RoomNumber: "G100", RoomType: "Lecture Hall", RoomCapacity: 100}
	b3room5 := model.Room{RoomNumber: "D180", RoomType: "Classroom", RoomCapacity: 25}
	b3room6 := model.Room{RoomNumber: "D180", RoomType: "Classroom", RoomCapacity: 25}
	b3room7 := model.Room{RoomNumber: "D180", RoomType: "Classroom", RoomCapacity: 30}
	b3room8 := model.Room{RoomNumber: "D180", RoomType: "Classroom", RoomCapacity: 30}
	b3room9 := model.Room{RoomNumber: "D180", RoomType: "Classroom", RoomCapacity: 35}
	b3room10 := model.Room{RoomNumber: "G190", RoomType: "Lecture Hall", RoomCapacity: 100}

	db.Create(&b2room1)
	db.Create(&b2room2)
	db.Create(&b2room3)
	db.Create(&b2room4)
	db.Create(&b2room5)
	db.Create(&b2room6)
	db.Create(&b2room7)
	db.Create(&b2room8)
	db.Create(&b2room9)
	db.Create(&b2room10)

	db.Create(&room1)
	db.Create(&room2)
	db.Create(&room3)
	db.Create(&room4)
	db.Create(&room5)
	db.Create(&room6)
	db.Create(&room7)
	db.Create(&room8)
	db.Create(&room9)
	db.Create(&room10)

	db.Create(&b3room1)
	db.Create(&b3room2)
	db.Create(&b3room3)
	db.Create(&b3room4)
	db.Create(&b3room5)
	db.Create(&b3room6)
	db.Create(&b3room7)
	db.Create(&b3room8)
	db.Create(&b3room9)
	db.Create(&b3room10)

	fmt.Println("Creating locations")

	//should be 30 locations because their is 3 buildings and 10 rooms in each.  3 * 10 = 30
	location1 := model.Location{BuildingID: building.BuildingID, RoomID: room1.RoomID}
	location2 := model.Location{BuildingID: building.BuildingID, RoomID: room2.RoomID}
	location3 := model.Location{BuildingID: building.BuildingID, RoomID: room3.RoomID}
	location4 := model.Location{BuildingID: building.BuildingID, RoomID: room4.RoomID}
	location5 := model.Location{BuildingID: building.BuildingID, RoomID: room5.RoomID}
	location6 := model.Location{BuildingID: building.BuildingID, RoomID: room6.RoomID}
	location7 := model.Location{BuildingID: building.BuildingID, RoomID: room7.RoomID}
	location8 := model.Location{BuildingID: building.BuildingID, RoomID: room8.RoomID}
	location9 := model.Location{BuildingID: building.BuildingID, RoomID: room9.RoomID}
	location10 := model.Location{BuildingID: building.BuildingID, RoomID: room10.RoomID}

	location11 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room1.RoomID}
	location12 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room2.RoomID}
	location13 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room3.RoomID}
	location14 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room4.RoomID}
	location15 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room5.RoomID}
	location16 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room6.RoomID}
	location17 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room7.RoomID}
	location18 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room8.RoomID}
	location19 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room9.RoomID}
	location20 := model.Location{BuildingID: building2.BuildingID, RoomID: b2room10.RoomID}

	location21 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room1.RoomID}
	location22 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room2.RoomID}
	location23 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room3.RoomID}
	location24 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room4.RoomID}
	location25 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room5.RoomID}
	location26 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room6.RoomID}
	location27 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room7.RoomID}
	location28 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room8.RoomID}
	location29 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room9.RoomID}
	location30 := model.Location{BuildingID: building2.BuildingID, RoomID: b3room10.RoomID}




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

	db.Create(&location11)
	db.Create(&location12)
	db.Create(&location13)
	db.Create(&location14)
	db.Create(&location15)
	db.Create(&location16)
	db.Create(&location17)
	db.Create(&location18)
	db.Create(&location19)
	db.Create(&location20)

	db.Create(&location20)
	db.Create(&location21)
	db.Create(&location22)
	db.Create(&location23)
	db.Create(&location24)
	db.Create(&location25)
	db.Create(&location26)
	db.Create(&location27)
	db.Create(&location28)
	db.Create(&location29)
	db.Create(&location30)

	day1 := model.Day{MeetingDay: "MW"}
	day2 := model.Day{MeetingDay: "TR"}
	db.Create(&day1)
	db.Create(&day2)

	semester1 := model.Semester{Year: 2018, Season: "Spring", SemesterStatus:"Closed"}
	semester2 := model.Semester{Year: 2018, Season: "Fall", SemesterStatus:"Closed"}
	winterSemester := model.Semester{Year:2018, Season:"Winter", SemesterStatus:"Closed"}
	summerSemester := model.Semester{Year:2018, Season:"Summer" , SemesterStatus:"Closed"}

	fallLastYear := model.Semester{Year:2017, Season:"Fall" , SemesterStatus:"Closed"}
	springLastYear := model.Semester{Year:2017, Season:"Spring" , SemesterStatus:"Closed"}
	winterLastYear := model.Semester{Year:2017, Season:"Winter" , SemesterStatus:"Closed"}
	summerLastYear := model.Semester{Year:2017, Season:"Summer" , SemesterStatus:"Closed"}

	fallLastLastYear := model.Semester{Year:2016, Season:"Fall" , SemesterStatus:"Closed"}
	springLastLastYear := model.Semester{Year:2016, Season:"Spring" , SemesterStatus:"Closed"}

	fallLastLastLastYear := model.Semester{Year:2015, Season:"Fall" , SemesterStatus:"Closed"}
	springLastLastLastYear := model.Semester{Year:2015, Season:"Spring" , SemesterStatus:"Closed"}

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


