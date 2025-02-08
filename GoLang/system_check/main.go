package main

import (
    "fmt"
    "os/exec"
    "strings"
    "time"
)

var c_white string = "\033[0m";
var c_cyan string = "\033[036m";
var c_red string = "\033[31m";
var c_yellow string = "\033[33m";

func header(title string) {
    fmt.Printf("%s------------------\n%6s\n------------------\n%s", c_cyan, title, c_white);
}

func getIO() {
    header("IO STATS");
    result, err := exec.Command("iostat", "1", "10").Output();
    if (err != nil) {
        fmt.Printf("%s", err);
    }
    output := string(result[:]);
    fmt.Println(output);
}

func getProcesses() {
    header("PROCESSES STUCK IN 'D' STATE");
    var output string;
    for i:=1; i<=30; i+=1 {
        result, err := exec.Command("ps", "-eo", "state,pid").Output();
        if (err != nil) {
            fmt.Printf("%s", err);
        }
        if (strings.HasPrefix(string(result),"D")) {
            output = string(result[:]);
        } else {
            output = "-";
        }
        time.Sleep(2);
        fmt.Println(output);
    }
}

func getCPU() {
    header("CPU USAGE");
    result, err := exec.Command("top", "-bn", "1").Output();
    if (err != nil) {
        fmt.Printf("%s", err);
    }
    output := string(result);
    fmt.Println(output);
}

func getNetConfig() {
    header("IP INFORMATION");
    result, err := exec.Command("ip", "-0", "addr").Output();
    if (err != nil) {
        fmt.Printf("%s", err);
    }
    output := string(result);
    fmt.Println(output);

    result, err = exec.Command("ip", "-0", "route").Output();
    if (err != nil) {
        fmt.Printf("%s", err);
    }
    output = string(result);
    fmt.Println(output);

    result, err = exec.Command("cat /etc/systemd/network/05-eth0.network").Output();
    if (err != nil) {
        fmt.Printf("%sERROR: File '/etc/systemd/network/05-eth0.network' not found!%s", c_red, c_white);
    }
    output = string(result);
    fmt.Println(output);
}

func main() {
    getIO();
    getProcesses();
    getCPU();
//    getMTR(ip_addr);
    getNetConfig();
}
