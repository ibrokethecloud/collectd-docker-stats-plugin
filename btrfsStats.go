// Function to return btrfs stats for each container //
package main

import (
    "fmt"
    //"github.com/docker/engine-api/client"
    "github.com/docker/engine-api/types"
    //"golang.org/x/net/context"
    //"flag" //importing flag package for command line flag parsing
    //"os"
    //"net/http"
)


func btrfsStats (containers []types.Container, modePtr *string){
	// Loop on each container and get the size of folder on filesystem // 
	// Dont intend to return anything yet. Just dump to SystemOut //
	var container types.Container

	for _, container = range containers{

		dirpath := "/var/lib/docker/containers/"+container.ID
    	dirsize,err := DirSizeMB(dirpath)
    	if err != nil {
        	panic(err)
    	} else {
    		fmt.Printf("%s - %d \n",container.Names,dirsize)
    	}
		
	}
	
	/*fsAllocation := make(map[string]int64)	
	

	fmt.Println(containers)
	fmt.Println(*modePtr) */
}