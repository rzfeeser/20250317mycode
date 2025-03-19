package main

import (
    "fmt"
)

type Hurricane struct {
    wind_speed int // hurricane wind speeds
    downfall float32 // hurricane wind amounts
    name string // hurricane name
}

// determines if hurricane is dangerous to the public
func (self Hurricane) isDangerous() bool {
    if self.wind_speed >= 90 {
        return true
    }
    return false
}



func main() {
    hurricane01 := Hurricane{90,4.8,"Albert"}
    hurricane02 := Hurricane{110,5.1,"Bobert"}
    hurricane03 := Hurricane{88,3.3,"Charlie"}
    hurricane04 := Hurricane{downfall: 6.7, name: "Zad", wind_speed: 45}

    fmt.Println(hurricane01, hurricane02)

    // init and declare
    hurricane_tracker := []Hurricane{hurricane01, hurricane02, hurricane03, hurricane04}

    // display hurricane03 from our slice (charlie)
    fmt.Println(hurricane_tracker[2])

    hurricane_tracker_backup := make([]Hurricane, 0)

    hurricane_tracker_backup = append(hurricane_tracker_backup, hurricane01, hurricane02)

    fmt.Println(hurricane_tracker_backup)


    // try out our new receiver function (method)
    fmt.Println(hurricane01.isDangerous())

    for _, storm := range hurricane_tracker {
        fmt.Println(storm.isDangerous())
    }

}
