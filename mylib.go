package my-library-go
import (
    "fmt"
    "bufio"
    "os"
    "time"
    "strings"
)
func InputSplit(s string) []string {
    scanner:=bufio.NewScanner(os.Stdin)
    scanner.Scan()
    str:=scanner.Text()
    parts:=strings.Split(str,s)
    return parts
}
func ParseTimeParts(l,p string) time.Time {
    parse,err:=time.Parse(l,p)
        if err!=nil {
            panic(err)
        }
    return parse
}
func MaxHour(a,b time.Time) time.Time {
    if a.Hour()>b.Hour() {
        return a
    } else {
        return b
    }
}
func CheckTime(a,b time.Time) time.Duration {
    diff:=a.Sub(b)
    if diff<0 {
        diff=-diff
    }
    return diff
}
