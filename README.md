tfg
===

A Python-Style time parse/format for Golang.

Golang's time format style "01/02 03:04:05PM '06 -0700" is toooooooo weird!(http://golang.org/pkg/time/#Constants)

Python's style is much easier.

See at:http://docs.python.org/2/library/datetime.html#strftime-strptime-behavior


usage:
===
VERY SIMPLE!

    package main
    
    import(
      "github.com/Codefor/tfg"
      "time"
    )
      
    func main(){
      //format
      tfg.Format(time.Now(),"%Y-%m-%d %H:%M:%S")
      //parse
      tfg.Parse("%Y-%m-%d %H:%M:%S","2013-12-01 01:00:35")
    }
