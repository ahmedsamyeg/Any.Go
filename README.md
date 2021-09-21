# Any.Go

A library to generate random values for variables and types members.

Usage:
```Go
package main

import (
	"fmt"

	"./Any"
)

func main() {
	fmt.Println("Random Country:", Any.Country())
	fmt.Println("Random Month:", Any.Month())
	fmt.Println("Random US State:", Any.UsState())
	fmt.Println("Random Solar System Planet:", Any.SolarSystemPlanet())
	fmt.Println("Random Day of the week:", Any.DayOftheWeek())
}
```

Output

```
Random Country: China
Random Month: March
Random US State: California
Random Solar System Planet: Venus
Random Day of the week: March
```
