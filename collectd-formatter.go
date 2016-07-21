// Will be using this to format the output to be compatible with collectd plugin //
package main

import (
    "fmt"
    "strings"
)


func collectdFormatter(containerName string, metricValue float64, metricName string){
    // Accepts the map built by the earlier step and formats it up //
    // Getting current Hostname which should be available from collectd plugin //
    //print("PUTVAL {}/exec-btrfs_{}/gauge-bytes_total interval={} N:{:.0f}".format(hostname, fs_name, interval, total))
    containerName = strings.TrimLeft(containerName,"/")
    //fmt.Println(containerName)
    fmt.Printf("PUTVAL %s/exec-%s/gauge-%s interval=%s N:%.00f\n",GLOBAL_HOSTNAME, containerName, metricName, GLOBAL_INTERVAL, metricValue)
} 