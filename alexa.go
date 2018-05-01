package main

import (
  // "encoding/json"
  // "errors"
  "log"
  "math/rand"
  "net/http"
  // "os"
  // "strconv"
  // "strings"
  "time"

  "gobot.io/x/gobot"
  "gobot.io/x/gobot/platforms/dji/tello"

  // "github.com/kennygrant/sanitize"
  alexa "github.com/mikeflynn/go-alexa/skillserver"
)

var Drone = "TODO"

var Applications = map[string]interface{}{
  "/echo/tails": alexa.EchoApplication{
    AppID:   "amzn1.ask.skill.91ab0818-a4f2-436f-8827-923f41226522",
    Handler: EchoTails,
  },
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  alexa.Run(Applications, "8080")
}

var TailsGreetings = []string{
  "Clear skies ahead!",
}

var TailsConfirmations = []string{
  "Roger that!",
  "Copy that!",
  "Affirmative!",
  "Sure thing!",
}

func EchoTails(w http.ResponseWriter, r *http.Request) {
  echoReq := alexa.GetEchoRequest(r)

  log.Println(echoReq.GetRequestType())
  log.Println(echoReq.GetSessionID())

  if echoReq.GetRequestType() == "LaunchRequest" {
    echoResp := tailsStart(echoReq)

    json, _ := echoResp.String()
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.Write(json)
  } else if echoReq.GetRequestType() == "IntentRequest" {
    log.Println(echoReq.GetIntentName())

    var echoResp *alexa.EchoResponse

    switch echoReq.GetIntentName() {
    case "TakeOff":
      echoResp = tailsTakeOff(echoReq)
    case "Land":
      echoResp = tailsLand(echoReq)
    case "ShutDown":
      echoResp = tailsStop(echoReq)
    default:
      echoResp = alexa.NewEchoResponse().OutputSpeech("Say what?").EndSession(false)
    }

    json, _ := echoResp.String()
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.Write(json)

  } else if echoReq.GetRequestType() == "SessionEndedRequest" {
    echoResp := tailsStop(echoReq)
    json, _ := echoResp.String()
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.Write(json)
  }
}

func tailsStart(echoReq *alexa.EchoRequest) (*alexa.EchoResponse) {
  msg := TailsGreetings[rand.Intn(len(TailsGreetings))]

  // Drone = tello.NewDriver("8888")

  echoResp := alexa.NewEchoResponse().OutputSpeech(msg).EndSession(false)

  return echoResp
}

func tailsStop(echoReq *alexa.EchoRequest) (*alexa.EchoResponse) {
  msg := "Uplink severed"

  // todo disconnect drone
  // Drone = ...

  echoResp := alexa.NewEchoResponse().OutputSpeech(msg).EndSession(true)

  return echoResp
}


func tailsTakeOff(echoReq *alexa.EchoRequest) (*alexa.EchoResponse) {
  msg := "Launch sequence initiated"

  // Drone.TakeOff()

  echoResp := alexa.NewEchoResponse().OutputSpeech(msg).EndSession(false)

  return echoResp
}


func tailsLand(echoReq *alexa.EchoRequest) (*alexa.EchoResponse) {
  msg := "Landing gear ready"

  // Drone.Land()

  echoResp := alexa.NewEchoResponse().OutputSpeech(msg).EndSession(false)

  return echoResp
}

// category, err := echoReq.GetSlotValue("Category")
//   _, catExists := JeopardyCategories[category]
//   if err != nil || !catExists {
//     catNames := []string{}
//     for k, _ := range JeopardyCategories {
//       catNames = append(catNames, k)
//     }

//     category = getRandom(catNames)

//     msg = msg + getRandom(JeopardyCatSelect) + category + ". "
//   } else {
//     category = strings.ToLower(category)
//   }



func getRandom(list []string) string {
  return list[rand.Intn(len(list))]
}

