// Examines struct composition
package main

import (
	"time"
	"fmt"
)

func main() {
	type City struct {
		Name string
		Population int
		Established time.Time
	}
	type State struct {
		Name string
		Formed time.Time
		Cities []City
	}
	type Country struct {
		Name string
		Formed time.Time
		States []State
	}
	pune := City{Name: "Pune", Population: 4000000}
	maha := State{Name: "Maharashtra", Cities: make([]City, 10)}
	maha.Cities[0] = pune
	india := Country {Name: "India"}
	india.States = append(india.States, maha)
	// increment pune's population
	india.States[0].Cities[0].Population += 1000
	fmt.Printf("%d\n", india.States[0].Cities[0].Population)

}
