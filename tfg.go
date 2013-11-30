package tfg

import (
	"time"
	"strings"
)

var m map[string]string

func init() {
	m = map[string]string{
		"%a": "Mon",//Short Week
		"%A": "Monday",//Week
		"%Y": "2006",//Year
		"%y": "06",//Short Year
		"%m": "01",//Month
		"%d": "02",//Day
		"%I": "03",//12-hour clock
		"%H": "15",//24-hour clock
		"%M": "04",//Minute
		"%S": "05",//Second
	}
}

func transform(in string)(out string){
	out = in
	for k, v := range m {
		if strings.Index(in,k) != -1 {
			out = strings.Replace(out, k, v, -1)
		}
	}
	return
}

func Parse(layout, value string) (time.Time, error) {
	layout = transform(layout)
	return time.Parse(layout,value)
}

func Format(t time.Time, layout string) string {
	layout = transform(layout)
	return t.Format(layout)
}

