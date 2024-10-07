package main

import (
    "encoding/json"
    "net/http"
    "os/exec"
    "strings"
)

type SystemInfo struct {
    IP        string `json:"ip"`
    Processes string `json:"processes"`
    DiskSpace string `json:"disk_space"`
    Uptime    string `json:"uptime"`
}

func getSystemInfo() SystemInfo {
    ip, _ := exec.Command("hostname", "-I").Output()
    processes, _ := exec.Command("ps", "-ax").Output()
    diskSpace, _ := exec.Command("df", "-h", "/").Output()
    uptime, _ := exec.Command("uptime", "-p").Output()

    return SystemInfo{
        IP:        strings.TrimSpace(string(ip)),
        Processes: strings.TrimSpace(string(processes)),
        DiskSpace: strings.TrimSpace(string(diskSpace)),
        Uptime:    strings.TrimSpace(string(uptime)),
    }
}

func systemInfoHandler(w http.ResponseWriter, r *http.Request) {
    info := getSystemInfo()
    json.NewEncoder(w).Encode(info)
}

func main() {
    http.HandleFunc("/", systemInfoHandler)
    http.ListenAndServe(":5000", nil)
}
