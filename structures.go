package main

import "fmt"

var forms = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

type human struct {
	name    string
	surname string
	age     int
	sex     string
	email   string
	phone   string
}

type worker struct {
	personInfo          *human
	education           string
	yearsOfExperience   int
	onVacation          bool
	company             string
	workingPost         string
	salaryPerHour       float32
	workingHoursPerWeek int
}

type schoolTeacher struct {
	basics     *worker
	Classes    []int  `json:"classes,omitempty"`
	Subject    string `json:"studyingSubject,omitempty"`
	ClassTrips bool   `json:"trips,omitempty"`
}

type programmer struct {
	basics              *worker
	Post                string   `json:"post,omitempty"`
	Languages           []string `json:"languages,omitempty"`
	FreelanceExperience bool     `json:"freelance,omitempty"`
}

type footballer struct {
	basics           *worker
	PlayPosition     string `json:"position"`
	League           string `json:"league,omitempty"`
	MonthsOfContract int    `json:"contract,omitempty"`
}

type doctor struct {
	basics          *worker
	License         string `json:"license"`
	Hospital        string `json:"hospital"`
	NeededEquipment string `json:"equipment,omitempty"`
}

type driver struct {
	basics                *worker
	DrivingCategories     string `json:"categories"`
	Region                string `json:"region,omitempty"`
	ForeignCountriesVisas string `json:"visas,omitempty"`
}

func (hum *human) getContact() {
	fmt.Printf("Contacts of %s %s: email - %s, phone - %s. \n", hum.name, hum.surname, hum.email, hum.phone)
}

func (wor *worker) getSalary() {
	fmt.Printf("%s %s earns %.2f$ per hour and works %d hours per week. \n", wor.personInfo.name, wor.personInfo.surname, wor.salaryPerHour, wor.workingHoursPerWeek)
}

func declaration(people []*worker) {
	fmt.Printf("We declared such people:\n")
	for _, v := range people {
		fmt.Printf("%s %s \n", v.personInfo.name, v.personInfo.surname)
	}
}

func contactsAndSalary(people []*worker) {
	for _, v := range people {
		v.personInfo.getContact()
		v.getSalary()
	}
}

func main() {

	var teacher1 = schoolTeacher{
		basics: &worker{
			personInfo: &human{
				name:    "Dale",
				surname: "Green",
				age:     27,
				sex:     "male",
				email:   "dalegreen@gmail.com",
				phone:   "123-45-67",
			},
			education:           "high",
			yearsOfExperience:   3,
			onVacation:          false,
			company:             "Government",
			workingPost:         "Teacher in high school",
			salaryPerHour:       15.0,
			workingHoursPerWeek: 8,
		},
		Classes:    forms[8:12],
		Subject:    "Science",
		ClassTrips: false,
	}

	var programmer1 = programmer{
		basics: &worker{
			personInfo: &human{
				name:    "Ann",
				surname: "Purple",
				age:     23,
				sex:     "female",
				email:   "ann_purple@gmail.com",
				phone:   "765-43-21",
			},
			education:           "high",
			yearsOfExperience:   2,
			onVacation:          false,
			company:             "Comp&CO",
			workingPost:         "Programmer",
			salaryPerHour:       15.0,
			workingHoursPerWeek: 8,
		},
		Post:                "Junior",
		Languages:           []string{"C#", "C++", "Java", "Python"},
		FreelanceExperience: true,
	}

	workers := make([]*worker, 0, 0)
	var t1 *schoolTeacher = &teacher1
	var p1 *programmer = &programmer1
	workers = append(workers, t1.basics, p1.basics)
	declaration(workers)
	contactsAndSalary(workers)

	// Как вызвать панику при помощи указателей - обратиться к полям объявленной, но не инициализированной структуры (через указатель)
	//var pPanic *schoolTeacher
	//fmt.Printf("We declared such people: " + pPanic.basics.personInfo.name + " " + pPanic.basics.personInfo.surname)

}
