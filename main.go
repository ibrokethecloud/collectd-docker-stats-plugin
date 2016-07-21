// Attempt to get Container stats using the go docker-engine apis //
package main

import (
    "fmt"
    "github.com/docker/engine-api/client"
    "github.com/docker/engine-api/types"
    "golang.org/x/net/context"
    "flag" //importing flag package for command line flag parsing
    "runtime"
    "os"
)

var GLOBAL_HOSTNAME string
var GLOBAL_INTERVAL string

func main() {
    
    runtime.GOMAXPROCS(2)

    // Defined Input flags for parsing //
    modePtr := flag.String("mode","collectd/introscope","String detailing what mode arguments are needed")
    statsPtr := flag.String("stats","btrfs/cpu/memory/containerfs","String detailing what stats are needed")
    connPtr := flag.String("connect","unix/tcp","String detailing what mode to use to connect to docker daemon")
    serverId := flag.String("serverid","Hostname","String detailing server hostname to be passed to script")
    interval := flag.String("interval","Interval at which to execute the check again","Wait for interval seconds before running script again")

    // Parsing the flags //
    flag.Parse()

    // Assigning values to the global variables //

    GLOBAL_HOSTNAME = *serverId
    GLOBAL_INTERVAL = *interval
    // Connect to docker daemon based on connection mode //
    var cli *client.Client
    var err error

    if *connPtr == "unix" {
        defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
        cli, err = client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
        //return cli, err
    } else if  *connPtr == "tcp" {
        // Get connection based on environment variables, which is used in case of TCP sockets
        cli, err = client.NewEnvClient()
        //return cli, err
    } else {
        fmt.Println("Undefined connection mode specified. Possible options are unix and tcp only")
        os.Exit(1)
    }
    
    if err != nil {
        panic(err)
    }

    // Set options for the query
    options := types.ContainerListOptions{All: true}
    
    // Get the list of containers 
    containers, err := cli.ContainerList(context.Background(), options)
    if err != nil {
        panic(err)
    }

    // Decide on which stats to gather now //
    if *statsPtr == "containerfs"{
        // Function call to find the biggest container //
        containerFsStats(containers,modePtr)
    } else if *statsPtr == "cpu" || *statsPtr == "mem"{
        // Function call to find container CPU stats
        containerUsageStats(cli,containers,statsPtr,modePtr)
    } else if *statsPtr == "btrfs"{
        // Function call to find container memory stats
        btrfsStats(modePtr)
    } else {
        fmt.Println("Request stat has not been implemented yet")
        os.Exit(1)
    }

    /* fmt.Println(containers)
    fmt.Println("\n\n")
    for _, c := range containers {
        fmt.Println(c)
    }

    dirpath := "/Users/gauravmehta/work/repos/"

    dirsize,err := DirSizeMB(dirpath)
    if err != nil {
        panic(err)
    }

    fmt.Printf("--- Size of folder %s is %d --- %s -- %s ", dirpath, dirsize, *modePtr, *statsPtr) */
}