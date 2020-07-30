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

func main() {
	var (
		teacher1 = schoolTeacher{
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
				salaryPerHour:       20.0,
				workingHoursPerWeek: 8,
			},
			Classes:    forms[8:12],
			Subject:    "Science",
			ClassTrips: false,
		}
	)
	var p *schoolTeacher = &teacher1

	fmt.Printf("We declared such people: " + p.basics.personInfo.name + " " + p.basics.personInfo.surname + "\n")
	teacher1.basics.personInfo.getContact()
	teacher1.basics.getSalary()

	// Как вызвать панику при помощи указателей - обратиться к полям объявленной, но не инициализированной структуры (через указатель)
	//var pPanic *schoolTeacher
	//fmt.Printf("We declared such people: " + pPanic.basics.personInfo.name + " " + pPanic.basics.personInfo.surname)

}
