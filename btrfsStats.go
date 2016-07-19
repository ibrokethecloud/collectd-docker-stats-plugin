// Function to return btrfs stats for each container //
package main

import (
    "fmt"
    "strings"
    //"golang.org/x/net/context"
    //"flag" //importing flag package for command line flag parsing
    "os"
    "os/exec"
    "bytes"
    "log"
    "strconv"
)


func btrfsStats (modePtr *string){

    // Query Docker Daemon to see if BTRFS is in use
    // If in use get BTRFS stats in a map and then format the output
    path, err := exec.LookPath("docker")
    if err != nil {
        log.Fatal("docker not found in lookup path")
        os.Exit(1)
    }
    //fmt.Printf("Docker is available at %s", path)
    
    // Going to query Docker Daemon for information on daemon //
    cmd := exec.Command(path, "info")
    var out bytes.Buffer
    cmd.Stdout = &out
    err = cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Printf("Docker config result: %q\n", out.String())

    //var dockerInfo map[string]string
    var dockerInfoStringArray,tmpString,btrfsInfoStringArray []string
    var allocated,used int64
    allocated, used = 0, 0

    dockerInfo := make(map[string]string)
    btrfsInfo := make(map[string]int64)
    dockerInfoStringArray = strings.Split(out.String(), "\n")
    
    //fmt.Println(dockerInfoStringArray[0])

     for _,pair := range dockerInfoStringArray {
        if len(pair) > 0 {
            tmpString = strings.Split(pair,":")
            //fmt.Printf("%s - %s \n", tmpString[0],tmpString[1])
            dockerInfo[tmpString[0]] = strings.TrimSpace(tmpString[1])
        }
    }

    //fmt.Println(dockerInfo["Storage Driver"])
    //fmt.Println("\n\n")

    if dockerInfo["Storage Driver"] == "btrfs" {
        //Then run the btrfs commands
        btrfspath := "/sbin/btrfs"
        btrfsargs := [] string{"fi", "df", "-b", "/var/lib/docker"}

        // Resetting the output buffer before re-use //

        out.Reset()
        cmd = exec.Command(btrfspath,btrfsargs...)
        cmd.Stdout = &out
        err = cmd.Run()

        if err != nil {
            log.Fatal(err)
            fmt.Println("btrfs execution failed")
        }
        btrfsInfoStringArray = strings.Split(out.String(),"\n")

        for _,pair := range btrfsInfoStringArray{
            if len(pair) > 0 {
                tmpString = strings.Split(pair,",")
                tmpAllocatedString := strings.Split(tmpString[1],"=")
                tmpUsedString := strings.Split(tmpString[2],"=")

                //debug statements
                //fmt.Println(tmpAllocatedString[1])
                //fmt.Println(tmpUsedString[1])

                tmpAllocated,errAllocated := strconv.ParseInt(tmpAllocatedString[1],10,64)
                tmpUsed,errUsed := strconv.ParseInt(tmpUsedString[1],10,64)

                if errAllocated != nil || errUsed != nil {
                    
                    log.Fatal(err)
                }
                
                allocated += tmpAllocated
                used += tmpUsed

            }

            btrfsInfo["allocated"] = allocated
            btrfsInfo["used"] = used
            
        }

        // Using DF to find total space in filesystem //
        dfpath,err := exec.LookPath("df")
        if err != nil{
            log.Fatal(err)
        }

        dfargs := [] string{"/var/lib/docker"}
        out.Reset()
        cmd = exec.Command(dfpath,dfargs...)
        cmd.Stdout = &out
        err = cmd.Run()
        
        if err != nil {
            log.Fatal(err)
            fmt.Println("df /var/lib/docker execution failed")
        }
        dfInfoArray := strings.Split(out.String(),"\n")

        tmpTotalString := strings.Fields(dfInfoArray[1])
        tmpTotal, errTotal := strconv.ParseInt(tmpTotalString[1],10,64)

        if errTotal != nil {
            log.Fatal(err)
            fmt.Println("total is invalid")
        }

        btrfsInfo["total"] = tmpTotal*1024
        
        fmt.Println(btrfsInfo)
        

    } else {
        log.Fatal("BTRFS is not being used by the docker daemon")
        os.Exit(1)
    }
}