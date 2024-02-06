package main

import (
  "fmt"
  "math/rand" 
  "time"

  MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

  // AWS IoT MQTT setup
  endpoint := "xxxxxxxxxxxxxx.iot.us-east-1.amazonaws.com"
  
  opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("ssl://%s:8883", endpoint))
  opts.SetClientID("myClient")
  
  // Connect and publish
  client := MQTT.NewClient(opts)
  if token := client.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }

  // Publish every 5 minutes
  ticker := time.NewTicker(5 * time.Minute)
  for range ticker.C {
    
    // Generate random temp & humidity
    temp := rand.Intn(50) 
    humi := rand.Intn(100)

    // Publish message
    payload := fmt.Sprintf(`{"temp": %d, "humi": %d}`, temp, humi)
    token := client.Publish("my/topic", 0, false, payload)
    token.Wait()
  }

}