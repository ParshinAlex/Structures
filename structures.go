package main

import "fmt"

var forms = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

// Structures

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

// Interface zone

type workerI interface {
	workPostContact()
}

func getWorkPostContact(w workerI) {
	w.workPostContact()
}

func (p schoolTeacher) workPostContact() {
	fmt.Printf("Worker %s %s: work - %s, subject - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.workingPost, p.Subject, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p programmer) workPostContact() {
	fmt.Printf("Worker %s %s: work - %s, post - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.workingPost, p.Post, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p footballer) workPostContact() {
	fmt.Printf("Worker %s %s: work - %s, position - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.workingPost, p.PlayPosition, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p doctor) workPostContact() {
	fmt.Printf("Worker %s %s: work - %s, license - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.workingPost, p.License, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p driver) workPostContact() {
	fmt.Printf("Worker %s %s: work - %s, region - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.workingPost, p.Region, p.basics.personInfo.email, p.basics.personInfo.phone)
}

// Cache zone

func printValuesOfCache(m map[string]workerI) {
	for _, v := range m {
		fmt.Printf("%T\n", v)
	}
}

// Main zone

func main() {

	// Variables

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
			workingHoursPerWeek: 40,
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
			workingHoursPerWeek: 40,
		},
		Post:                "Junior",
		Languages:           []string{"C#", "C++", "Java", "Python"},
		FreelanceExperience: true,
	}

	var footballer1 = footballer{
		basics: &worker{
			personInfo: &human{
				name:    "Rick",
				surname: "Black",
				age:     19,
				sex:     "male",
				email:   "rick_black@gmail.com",
				phone:   "111-22-33",
			},
			education:           "high school",
			yearsOfExperience:   1,
			onVacation:          true,
			company:             "Atletico",
			workingPost:         "Reserve Footballer",
			salaryPerHour:       10,
			workingHoursPerWeek: 20,
		},
		PlayPosition:     "Defender",
		League:           "Junior",
		MonthsOfContract: 6,
	}

	// Objects

	var t1 *schoolTeacher = &teacher1
	var p1 *programmer = &programmer1
	var f1 *footballer = &footballer1

	// Creating and using cache

	workers := make(map[string]workerI)
	workers["teacher1"] = t1
	workers["programmer1"] = p1
	workers["footballer1"] = f1

	for _, v := range workers {
		getWorkPostContact(v)
	}

	printValuesOfCache(workers)

}
