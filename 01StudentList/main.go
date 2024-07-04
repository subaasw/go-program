package main

import "fmt"

type StudentList []Student

type Student struct {
	name    string
	age     uint
	faculty string
	rollNo  uint
	gpa     float32
}

func (s Student) getName() string {
	return s.name
}

func (s Student) getAge() uint {
	return s.age
}

func (s Student) getFaculty() string {
	return s.faculty
}

func (s Student) getRollNo() uint {
	return s.rollNo
}

func (s Student) getGPA() float32 {
	return s.gpa
}

func (s Student) toString() string {
	return fmt.Sprintf("Student name is %s, age was %d, roll no %d, and Faculty is %s, GPA %f ", s.getName(), s.getAge(), s.getRollNo(), s.getFaculty(), s.getGPA())
}

// func (s *Student) setName(name string) {
// 	s.name = name
// }

func (s *Student) setAge(age uint) {
	s.age = age
}

// func (s *Student) setFaculty(faculty string) {
// 	s.faculty = faculty
// }

// func (s *Student) setRollNo(rollNo uint) {
// 	s.rollNo = rollNo
// }

func (s *Student) setGPA(gpa float32) {
	s.gpa = gpa
}

func printMenu() {
	fmt.Println()
	fmt.Print(`
Choose one option:
1. Add New student
2. remove a student
3. update student
4. display all students
5. exit
`)
	fmt.Println()
}

func setInput(message string, val interface{}) {
	fmt.Print(message)
	fmt.Scanln(val)
}

func addStudent(sl *StudentList) {
	var name string
	var age uint
	var gpa float32
	var faculty string
	var rollNo uint

	setInput("Enter Name: ", &name)
	setInput("Enter Age: ", &age)
	setInput("Enter Faculty: ", &faculty)
	setInput("Enter Roll No: ", &rollNo)
	setInput("Enter GPA: ", &gpa)

	student := Student{
		name:    name,
		age:     age,
		gpa:     gpa,
		rollNo:  rollNo,
		faculty: faculty,
	}

	*sl = append(*sl, student)
	// fmt.Println("\n" + student.toString())
	fmt.Println("\nStudent", name, "Successfully added")
}

func (sL StudentList) getUserByRollNo(rollNo uint) (Student, int, bool) {
	for index, student := range sL {
		if student.rollNo == rollNo {
			return student, index, true
		}

	}
	return sL[0], 0, false
}

func removeStudent(sL *StudentList) {
	var rollNo uint

	setInput("Delete user by Roll No: ", &rollNo)

	_, index, ok := sL.getUserByRollNo(rollNo)

	if ok {
		*sL = append((*sL)[:index], (*sL)[index+1:]...)
	}

}

func updateStudent(sL *StudentList) {
	var rollNo uint
	setInput("Enter the user Roll No: ", &rollNo)

	student, index, ok := sL.getUserByRollNo(rollNo)

	if ok {
		var age uint
		var gpa float32

		setInput("Enter Age: ", &age)
		setInput("Enter GPA: ", &gpa)

		student.setAge(age)
		student.setGPA(gpa)

		(*sL)[index] = student
	}
}

func main() {
	var choose int
	studentList := StudentList{}

	for {
		printMenu()

		fmt.Print("Select one: ")
		fmt.Scan(&choose)

		if choose == 1 {
			addStudent(&studentList)
		} else if choose == 2 {
			removeStudent(&studentList)
		} else if choose == 3 {
			updateStudent(&studentList)
		} else if choose == 4 {
			fmt.Println("\nStudents List are:")

			for _, student := range studentList {
				fmt.Println(student.toString())
			}
		} else if choose == 5 {
			fmt.Println("\n👋 Sayonara!")
			break
		} else {
			fmt.Println("\nSorry Wrong input!")
		}
	}
}
