package main

import  (
    "fmt"
)

type date struct {
    day int
    month int
    year int
}

func daysInYear(year int) int {
    if year % 400 == 0 || (year % 100 != 0 && year % 4 == 0) {
        return 366
    }
    return 365
}

func daysSince1900(year int) int {
    sum := 1
    for i := 1900; i < year; i++ {
        sum += daysInYear(i)
    }
    return sum
}

func dateFromDays(days int) date {
    months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}    
    monthName := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
    year := 1900
    for ; days > daysInYear(year); year++ {
        days -= daysInYear(year)
    }
    month := 0
    for ; days > months[month]; month++ {
        days -= months[month]
    }
    fmt.Printf("Date: %d, %s, %d\n", days, monthName[month], year)

    return date{days, month, year}
}

func sundaysOnFirstOfTheMonth(start date, end date) int {
    months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    count := daysSince1900(start.year)
    sundays := 0
    for year := start.year; year <= end.year; year++ {
        for month := 0; month < 12; month++ {
            // dateFromDays(count)
            if count % 7 == 0 {
                sundays++
            }
            if daysInYear(year) == 366 && month == 1 {
                count += 29
            } else {
                count += months[month]
            }
        }
    }
    return sundays
}

func main() {
    start := date{1, 0, 1901}
    end   := date{31, 11, 2000}
    fmt.Printf("Sundays landing on the first of the month: %d\n", 
                sundaysOnFirstOfTheMonth(start, end))
}