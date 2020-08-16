package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/rs/xid"
)

var forms = [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

// Safe Map

type safeMap struct {
	v   map[xid.ID]workerI
	mux sync.RWMutex
}

func (mapa *safeMap) Update(id xid.ID, value workerI) {
	mapa.mux.Lock()
	defer mapa.mux.Unlock()
	mapa.v[id] = value
}

// Function for printing both workers and heads, using waitgroups

func printAll(workers *safeMap, heads *safeMap) {
	var wg sync.WaitGroup

	wg.Add(len(heads.v))
	for _, value := range heads.v {
		go getWorkPostContact(value, &wg)
	}
	wg.Wait()

	wg.Add(len(workers.v))
	for _, value := range workers.v {
		go getWorkPostContact(value, &wg)
	}
	wg.Wait()
}

// Structures

type human struct {
	id      xid.ID
	name    string
	surname string
	age     string
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

func getWorkPostContact(w workerI, wg *sync.WaitGroup) {
	w.workPostContact()
	wg.Done()
}

func (p schoolTeacher) workPostContact() {
	fmt.Printf("Worker %s %s: id - %s, work - %s, subject - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.personInfo.id, p.basics.workingPost, p.Subject, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p programmer) workPostContact() {
	fmt.Printf("Worker %s %s: id - %s, work - %s, post - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.personInfo.id, p.basics.workingPost, p.Post, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p footballer) workPostContact() {
	fmt.Printf("Worker %s %s: id - %s, work - %s, position - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.personInfo.id, p.basics.workingPost, p.PlayPosition, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p doctor) workPostContact() {
	fmt.Printf("Worker %s %s: id - %s, work - %s, license - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.personInfo.id, p.basics.workingPost, p.License, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p driver) workPostContact() {
	fmt.Printf("Worker %s %s: id - %s, work - %s, region - %s, contacts: email - %s, phone - %s. \n", p.basics.personInfo.name,
		p.basics.personInfo.surname, p.basics.personInfo.id, p.basics.workingPost, p.Region, p.basics.personInfo.email, p.basics.personInfo.phone)
}

func (p worker) workPostContact() {
	fmt.Printf("Worker %s %s: id - %s, work - %s, education - %s, years of experience - %d, company - %s, contacts: email - %s, phone - %s. \n",
		p.personInfo.name, p.personInfo.surname, p.personInfo.id, p.workingPost, p.education, p.yearsOfExperience, p.company,
		p.personInfo.email, p.personInfo.phone)
}

func (p human) workPostContact() {
	fmt.Printf("Human %s %s: id - %s, age - %s, sex - %s, contacts: email - %s, phone - %s. \n",
		p.name, p.surname, p.id, p.age, p.sex, p.email, p.phone)
}

// Cache zone

func printValuesOfCache(m map[string]workerI) {
	for _, v := range m {
		fmt.Printf("%T\n", v)
	}
}

// Type assertion zone

func workerToHuman(p workerI) human {
	t, ok := p.(human)
	if !ok {
		fmt.Printf("Error in type assertion! \n")
		return human{}
	}
	return t

}

// Main zone

func main() {

	// Variables

	var teacher1 = schoolTeacher{
		basics: &worker{
			personInfo: &human{
				id:      xid.New(),
				name:    "Dale",
				surname: "Green",
				age:     "27",
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
				id:      xid.New(),
				name:    "Ann",
				surname: "Purple",
				age:     "23",
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
				id:      xid.New(),
				name:    "Rick",
				surname: "Black",
				age:     "19",
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

	var doctor1 = doctor{
		basics: &worker{
			personInfo: &human{
				id:      xid.New(),
				name:    "Nick",
				surname: "Yellow",
				age:     "35",
				sex:     "male",
				email:   "nick_yellow@gmail.com",
				phone:   "123-22-33",
			},
			education:           "high",
			yearsOfExperience:   10,
			onVacation:          false,
			company:             "Government",
			workingPost:         "Surgeon",
			salaryPerHour:       20,
			workingHoursPerWeek: 20,
		},
		License:         "G2-50",
		Hospital:        "National Minnesote Hospital",
		NeededEquipment: "Surgery",
	}

	// Creating and using cache

	workers := safeMap{v: make(map[xid.ID]workerI)}
	heads := safeMap{v: make(map[xid.ID]workerI)}
	workers.Update(teacher1.basics.personInfo.id, teacher1)
	workers.Update(programmer1.basics.personInfo.id, programmer1)
	heads.Update(footballer1.basics.personInfo.id, footballer1)
	heads.Update(doctor1.basics.personInfo.id, doctor1)
	printAll(&workers, &heads)

	// Web Server

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "postform.html")
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {

		var person1 = human{
			id:      xid.New(),
			name:    r.FormValue("name"),
			surname: r.FormValue("surname"),
			age:     r.FormValue("age"),
			sex:     r.FormValue("sex"),
			email:   r.FormValue("email"),
			phone:   r.FormValue("phone"),
		}

		fmt.Fprintf(w, "Structure is created and uploaded to cache.\nIt's id: %s \nHuman %s %s: age - %s, sex - %s, contacts: email - %s, phone - %s. \n",
			person1.id, person1.name, person1.surname, person1.age, person1.sex, person1.email, person1.phone)
		workers.Update(person1.id, person1)
		printAll(&workers, &heads)
	})

	http.HandleFunc("/getform", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "getform.html")
		query := r.URL.Query()
		requestedID, err := xid.FromString(query.Get("id"))
		if err == nil {
			str, ok := workers.v[requestedID]
			if ok {
				fmt.Printf("Requested structure: \n")
				str.workPostContact()
			} else {
				fmt.Printf("Requested structure: \n")
				heads.v[requestedID].workPostContact()
			}
		}
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}
