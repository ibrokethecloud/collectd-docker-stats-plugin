// Will be using this to format the output to be compatible with introscope plugin //
package main

import (
    "fmt"
    "strings"
    "strconv"
    "net/http"
    "bytes"
    "io/ioutil"
)


func introscopeFormatter(containerName string, metricValue float64, metricName string){
    // Sample url for calling epagent curl -H "Content-Type: application/json" --data @epagent_rest_sample.json http://centos64rt2.ca.com:8080/apm/metricFeed
    containerName = strings.TrimLeft(containerName,"/")
    stringmetricValue := strconv.FormatFloat(metricValue,'f',0,64)
    remoteUrl := "http://localhost:8080/apm/metricFeed"

    metricString := "{\"metrics\":[{type:\"IntCounter\",name:\"DockerStats|" + containerName + ":" + metricName + "\",value:\"" + stringmetricValue + "\"}]}"
    jsonStr := []byte(metricString)
    req, err := http.NewRequest("POST", remoteUrl, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil{
      panic(err)
    }

    defer resp.Body.Close()
    fmt.Println("response status:", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    // Now need to pump this to a rest api for EPAgent
}
