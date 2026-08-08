package partyrobot

import "strconv"
// Welcome greets a person by name.
func Welcome(name string) string {
    return "Welcome to my party, "+name+"!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday "+name+"! You are now "+strconv.Itoa(age)+" years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    var notable string
    if table > 99 {
        notable = strconv.FormatInt(int64(table), 10)
    } else {
    	if table > 9 {
            notable = "0"+strconv.FormatInt(int64(table), 10)
        } else {
        	notable = "00"+strconv.FormatInt(int64(table), 10)
        }
    }
	return "Welcome to my party, "+name+"!\nYou have been assigned to table "+notable+". Your table is "+direction+", exactly "+strconv.FormatFloat(distance, 'f', 1, 64)+" meters from here.\nYou will be sitting next to "+neighbor+"."
}
