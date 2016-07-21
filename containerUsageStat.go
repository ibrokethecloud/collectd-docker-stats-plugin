package main

import (
    //"fmt"
    "strings"
    "github.com/docker/engine-api/types"
    "github.com/docker/engine-api/client"
    "encoding/json"
    "io/ioutil"
    "golang.org/x/net/context"
    "log"
    //"sync"
    //"runtime"
)


func containerUsageStats (cli *client.Client, containers []types.Container, statsPtr *string,modePtr *string){
	// Pass the client connection from the main function as don't need to redo the logic to see if we are using unix sockets or tcp connection

	
	var container types.Container
	// Iterate over all available containers in the map and identify running containers first //
	for _, container = range containers{

		//tmpContainerArray := strings.Fields(pair)
		//fmt.Printf("%s - %s - %s \n",container.ID, container.State, container.Status)
		// Check if container.Status contains the String running and then get stats for that container //
		// Check for metric type first and then if container is up before proceeding 
        var containerName string
        for _, containerName = range container.Names{

        }

		if *statsPtr == "cpu" {
			if strings.Contains(container.Status, "Up") == true{
				GenerateCPUStats(cli,container.ID,containerName,modePtr)
			}
		} else if *statsPtr == "mem" {
			if strings.Contains(container.Status, "Up") == true{
				GenerateMemStats(cli,container.ID,containerName,modePtr)
			}

		} else {
			log.Fatal("This functionality doesn't exist yet")
		}

		//fmt.Printf("%f \n", <-channel)
			
	}
}


func GenerateCPUStats (cli *client.Client, container_id string, containerName string,modePtr *string) {
	// New function to get CPU usage based on container ID passed to the call
		body, err := cli.ContainerStats(context.Background(), container_id, false)
		if err != nil {
			log.Fatal(err)
		}
		defer body.Close()
		content, err := ioutil.ReadAll(body)
		var dockerStats types.Stats
		err = json.Unmarshal(content, &dockerStats)
		cpuPercentage := 0.0
		cpuDelta := (dockerStats.CPUStats.CPUUsage.TotalUsage - dockerStats.PreCPUStats.CPUUsage.TotalUsage)
		systemDelta := (dockerStats.CPUStats.SystemUsage - dockerStats.PreCPUStats.SystemUsage)

		//fmt.Printf("%d - %d \n",cpuDelta,systemDelta)

		if systemDelta > 0.0 &&  cpuDelta > 0.0 {
			cpuPercentage = float64(cpuDelta) / float64(systemDelta) * float64(len(dockerStats.CPUStats.CPUUsage.PercpuUsage)) * 100.0
		}
		

			
	//fmt.Printf("%s - %f \n",container_id,cpuPercentage)
	if *modePtr == "collectd"{
		collectdFormatter(containerName, cpuPercentage, "cpuPercentage")
		} else if *modePtr == "introscope"{
			log.Fatal("Feature coming soon")
		}else {
			log.Fatal("Logging mode not yet available")
		}
	
	//return cpuPercentage
}

func GenerateMemStats (cli *client.Client, container_id string,containerName string,modePtr *string){
		body, err := cli.ContainerStats(context.Background(), container_id, false)
		if err != nil {
			log.Fatal(err)
		}
		defer body.Close()
		content, err := ioutil.ReadAll(body)
		var dockerStats types.Stats
		err = json.Unmarshal(content, &dockerStats)
		memPercentage := 0.0
	
		if dockerStats.MemoryStats.Limit != 0 {
				memPercentage = float64(dockerStats.MemoryStats.Usage) / float64(dockerStats.MemoryStats.Limit) * 100.0
			}
		//fmt.Printf("%d - %d \n",cpuDelta,systemDelta)
		if *modePtr == "collectd"{
			collectdFormatter(containerName, memPercentage, "memPercentage")
		} else if *modePtr == "introscope"{
			log.Fatal("Feature coming soon")
		}else {
			log.Fatal("Logging mode not yet available")
		}
		
	//fmt.Printf("%s - %f \n",container_id,memPercentage)
	
}