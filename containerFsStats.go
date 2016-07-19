// Function to return btrfs stats for each container //
package main

import (
    "fmt"
    "strings"
    "github.com/docker/engine-api/types"
    //"golang.org/x/net/context"
    //"flag" //importing flag package for command line flag parsing
    //"os"
    //"net/http"
)


func containerFsStats (containers []types.Container, modePtr *string){
	// Loop on each container and get the size of folder on filesystem // 
	// Dont intend to return anything yet. Just dump to SystemOut //
	var container types.Container

	for _, container = range containers{

		dirpath := "/var/lib/docker/containers/"+container.ID
    	dirsize,err := DirSizeMB(dirpath)
        var containerName string
        for _, containerName = range container.Names{

        }
    	if err != nil {
        	panic(err)
    	} else {
    		fmt.Printf("%s - %d \n",strings.Trim(containerName,"/"),dirsize)
    	}
		
	}
	
	/*fsAllocation := make(map[string]int64)	
	

	fmt.Println(containers)
	fmt.Println(*modePtr) */
}