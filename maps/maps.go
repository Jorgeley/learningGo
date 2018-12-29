package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	//another simple way:
	var mm = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(mm)

	//another simple way:
	var mmm = map[string]Vertex{
		"Bell Labs": { //we can ommit the top level type here
			40.68433, -74.39967,
		},
		"Google": { //we can ommit the top level type here
			37.42202, -122.08408,
		},
	}
	//adding more elements:
	mmm["IBM"] = Vertex{
			37.42202, -122.08408,
		}
	fmt.Println(mmm)
	//we can delete values in the map:
	delete(mmm, "IBM")
	fmt.Println(mmm)
	//this can be used to test if an element exists in the map:
	elem, exists := mmm["IBM"]
	fmt.Println(elem, exists)
}

