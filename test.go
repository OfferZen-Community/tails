package main

import (
        // "fmt"
        // "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
        drone := tello.NewDriver("8888")

        work := func() {
                drone.TakeOff()

                // gobot.After(1*time.Second, func() {
                        drone.Land()
                // })
        }

        robot := gobot.NewRobot("tello",
                []gobot.Connection{},
                []gobot.Device{drone},
                work,
        )

        robot.Start()
}
