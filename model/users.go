package model

import (
	"fmt"
)

// Users Struct is the model for the app
type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

var users = []Users{
	Users{
		ID: 1, FirstName: "Timothy", LastName: "Cox", Email: "tcox0@dion.ne.jp", Gender: "Male",
	},
	Users{
		ID: 2, FirstName: "Sean", LastName: "Medina", Email: "smedina1@addthis.com", Gender: "Male",
	},
	Users{
		ID: 3, FirstName: "Jonathan", LastName: "Tucker", Email: "jtucker2@tripadvisor.com", Gender: "Male",
	},
	Users{
		ID: 4, FirstName: "Donna", LastName: "Payne", Email: "dpayne3@cdbaby.com", Gender: "Female",
	},
	Users{
		ID: 5, FirstName: "Emily", LastName: "Elliott", Email: "eelliott4@pen.io", Gender: "Female",
	},
	Users{
		ID: 6, FirstName: "Howard", LastName: "Wallace", Email: "hwallace5@latimes.com", Gender: "Male",
	},
	Users{
		ID: 7, FirstName: "Jacqueline", LastName: "George", Email: "jgeorge6@soup.io", Gender: "Female",
	},
	Users{
		ID: 8, FirstName: "Heather", LastName: "Kelly", Email: "hkelly7@hubpages.com", Gender: "Female",
	},
	Users{
		ID: 9, FirstName: "Kenneth", LastName: "Fisher", Email: "kfisher8@wunderground.com", Gender: "Male",
	},
	Users{
		ID: 10, FirstName: "Joan", LastName: "Flores", Email: "jflores9@icq.com", Gender: "Female",
	},
	Users{
		ID: 11, FirstName: "Wanda", LastName: "Hunt", Email: "whunta@linkedin.com", Gender: "Female",
	},
	Users{
		ID: 12, FirstName: "Alice", LastName: "Gordon", Email: "agordonb@mozilla.com", Gender: "Female",
	},
	Users{
		ID: 13, FirstName: "Nicholas", LastName: "Harrison", Email: "nharrisonc@fastcompany.com", Gender: "Male",
	},
	Users{
		ID: 14, FirstName: "Timothy", LastName: "Hudson", Email: "thudsond@bloomberg.com", Gender: "Male",
	},
	Users{
		ID: 15, FirstName: "Nancy", LastName: "Lynch", Email: "nlynche@163.com", Gender: "Female",
	},
	Users{
		ID: 16, FirstName: "Ryan", LastName: "Stevens", Email: "rstevensf@dedecms.com", Gender: "Male",
	},
	Users{
		ID: 17, FirstName: "Shawn", LastName: "Little", Email: "slittleg@cnet.com", Gender: "Male",
	},
	Users{
		ID: 18, FirstName: "Frances", LastName: "Garrett", Email: "fgarretth@cargocollective.com", Gender: "Female",
	},
	Users{
		ID: 19, FirstName: "Brian", LastName: "Nelson", Email: "bnelsoni@eepurl.com", Gender: "Male",
	},
	Users{
		ID: 20, FirstName: "Harry", LastName: "Anderson", Email: "handersonj@about.com", Gender: "Male",
	},
	Users{
		ID: 21, FirstName: "Michelle", LastName: "Nelson", Email: "mnelsonk@tiny.cc", Gender: "Female",
	},
	Users{
		ID: 22, FirstName: "Scott", LastName: "Palmer", Email: "spalmerl@canalblog.com", Gender: "Male",
	},
	Users{
		ID: 23, FirstName: "Helen", LastName: "Day", Email: "hdaym@geocities.com", Gender: "Female",
	},
	Users{
		ID: 24, FirstName: "Aaron", LastName: "Torres", Email: "atorresn@rediff.com", Gender: "Male",
	},
	Users{
		ID: 25, FirstName: "Cheryl", LastName: "Morris", Email: "cmorriso@theglobeandmail.com", Gender: "Female",
	},
	Users{
		ID: 26, FirstName: "Heather", LastName: "Sims", Email: "hsimsp@about.me", Gender: "Female",
	},
	Users{
		ID: 27, FirstName: "Andrew", LastName: "Morales", Email: "amoralesq@cbsnews.com", Gender: "Male",
	},
	Users{
		ID: 28, FirstName: "Kevin", LastName: "Lane", Email: "klaner@epa.gov", Gender: "Male",
	},
	Users{
		ID: 29, FirstName: "Karen", LastName: "Perkins", Email: "kperkinss@geocities.com", Gender: "Female",
	},
	Users{
		ID: 30, FirstName: "Jane", LastName: "Jackson", Email: "jjacksont@icq.com", Gender: "Female",
	},
	Users{
		ID: 31, FirstName: "Roy", LastName: "Green", Email: "rgreenu@csmonitor.com", Gender: "Male",
	},
	Users{
		ID: 32, FirstName: "Louis", LastName: "Berry", Email: "lberryv@so-net.ne.jp", Gender: "Male",
	},
	Users{
		ID: 33, FirstName: "Donald", LastName: "Kennedy", Email: "dkennedyw@umich.edu", Gender: "Male",
	},
	Users{
		ID: 34, FirstName: "Edward", LastName: "Schmidt", Email: "eschmidtx@seattletimes.com", Gender: "Male",
	},
	Users{
		ID: 35, FirstName: "Brenda", LastName: "Bennett", Email: "bbennetty@cargocollective.com", Gender: "Female",
	},
	Users{
		ID: 36, FirstName: "Bonnie", LastName: "Carr", Email: "bcarrz@desdev.cn", Gender: "Female",
	},
	Users{
		ID: 37, FirstName: "Tammy", LastName: "Bailey", Email: "tbailey10@technorati.com", Gender: "Female",
	},
	Users{
		ID: 38, FirstName: "Peter", LastName: "Murray", Email: "pmurray11@mozilla.com", Gender: "Male",
	},
	Users{
		ID: 39, FirstName: "Kathryn", LastName: "Peterson", Email: "kpeterson12@bandcamp.com", Gender: "Female",
	},
	Users{
		ID: 40, FirstName: "Linda", LastName: "Carter", Email: "lcarter13@redcross.org", Gender: "Female",
	},
	Users{
		ID: 41, FirstName: "Scott", LastName: "Howell", Email: "showell14@sogou.com", Gender: "Male",
	},
	Users{
		ID: 42, FirstName: "Lillian", LastName: "Nichols", Email: "lnichols15@a8.net", Gender: "Female",
	},
	Users{
		ID: 43, FirstName: "Frank", LastName: "Wells", Email: "fwells16@google.es", Gender: "Male",
	},
	Users{
		ID: 44, FirstName: "Jean", LastName: "Wheeler", Email: "jwheeler17@mlb.com", Gender: "Female",
	},
	Users{
		ID: 45, FirstName: "Phyllis", LastName: "Arnold", Email: "parnold18@deliciousdays.com", Gender: "Female",
	},
	Users{
		ID: 46, FirstName: "Irene", LastName: "Mills", Email: "imills19@nhs.uk", Gender: "Female",
	},
	Users{
		ID: 47, FirstName: "Rose", LastName: "Anderson", Email: "randerson1a@google.pl", Gender: "Female",
	},
	Users{
		ID: 48, FirstName: "Harry", LastName: "Little", Email: "hlittle1b@google.fr", Gender: "Male",
	},
	Users{
		ID: 49, FirstName: "Ruby", LastName: "Rogers", Email: "rrogers1c@digg.com", Gender: "Female",
	},
	Users{
		ID: 50, FirstName: "Jonathan", LastName: "Carpenter", Email: "jcarpenter1d@furl.net", Gender: "Male",
	},
}

// GetUsers returns a list of users for the index route
func GetUsers() []Users {
	return users
}

// GetUserByID returns a user by the given id
func GetUserByID(ID int) (*Users, error) {
	for _, user := range users {
		if user.ID == ID {
			return &user, nil
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("User with ID %v not found", ID))
}
