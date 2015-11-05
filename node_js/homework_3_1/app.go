package main

import (
	"flag"
	"fmt"
	//"regexp"
	//"strings"

	// // "github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "trueanthem.com/nab/cmd/webd/views/control_panel"
	// "trueanthem.com/nab/config"
	// "trueanthem.com/nab/config/mongodb"
	// "trueanthem.com/nab/crawl"
	// "trueanthem.com/nab/database"
	// "trueanthem.com/nab/model"
	// "trueanthem.com/nab/scheduler"
)

func main() {
	// flag.Parse()
	// config.PrepareViper("Nab")

	session, err := mgo.Dial(mongodb.URI(nil))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// var x string
	sites, _ := sites(session)

	for _, site := range sites {
		if site.CaseSensitive != false || site.PrefersOpenGraphURL != false || site.FirstSentence != nil {
			fmt.Println("*", "-----------------------------")
			fmt.Println("site: ", site.Host)
			fmt.Println("CaseSensitive: ", site.CaseSensitive)
			fmt.Println("OG URL: ", site.PrefersOpenGraphURL)
			fmt.Println("FirstSentence: ", site.FirstSentence)
			fmt.Println("PathConstraints: ", site.PathConstraints)
			fmt.Println("Meta: ", site.PageMetaMapConf)
		}

	}
}

// func containsAll(aSlice []string, bSlice []string) bool {
// 	for _, s := range bSlice {
// 		if !contains(aSlice, s) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func contains(slice []string, item string) bool {
// 	set := make(map[string]struct{}, len(slice))
// 	for _, s := range slice {
// 		set[s] = struct{}{}
// 	}

// 	_, ok := set[item]
// 	return ok
// }

func clients(session *mgo.Session) {

	// ([]model.Client, database.Clients)
	// var clients []model.Client
	// clientsDB := database.NewClients(session)
	// clientsDB.Find(nil).All(&clients)
	// return clients, clientsDB
}

// func sites(session *mgo.Session) ([]crawl.Site, database.Sites) {
// 	sitesDB := database.NewSites(session)
// 	sites := []crawl.Site{}
// 	sitesDB.Find(nil).All(&sites)
// 	return sites, sitesDB
// }

// func schedules(session *mgo.Session) ([]scheduler.Schedule, *database.Schedules) {
// 	schedulesDB := database.NewSchedules(session)
// 	schedules := []scheduler.Schedule{}
// 	schedulesDB.Find(nil).Sort("name").All(&schedules)
// 	return schedules, schedulesDB
// }

// func posttypes(session *mgo.Session) ([]controlpanel.PostTypeView, *database.PostTypes) {
// 	postTypeDB := database.NewPostTypes(session)
// 	posttypes := []controlpanel.PostTypeView{}
// 	postTypeDB.Find(nil).Sort("name").All(&posttypes)
// 	return posttypes, postTypeDB
// }

// func accounts(session *mgo.Session) ([]model.Account, database.Accounts) {
// 	accountsDB := database.NewAccounts(session)
// 	accounts := []model.Account{}
// 	accountsDB.Find(nil).Sort("name").All(&accounts)
// 	return accounts, accountsDB
// }
