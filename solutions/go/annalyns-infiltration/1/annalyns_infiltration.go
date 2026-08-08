package annalyn

import "fmt"
// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	// panic("Please implement the CanFastAttack() function")
    return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	// panic("Please implement the CanSpy() function")
   return (knightIsAwake || archerIsAwake || prisonerIsAwake)
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	// panic("Please implement the CanSignalPrisoner() function")
    if (prisonerIsAwake) {
        return (!archerIsAwake && prisonerIsAwake)
    } else {
    	return prisonerIsAwake
    }
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	// panic("Please implement the CanFreePrisoner() function")
    if (petDogIsPresent){
        return !archerIsAwake
    } else {
    	return !knightIsAwake && !archerIsAwake && prisonerIsAwake
    }
}

func main(){
    var knightIsAwake = true
    var archerIsAwake = true
    var prisonerIsAwake = false
    var petDogIsPresent = false
	fmt.Println(CanFastAttack(knightIsAwake))
    fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
    fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
    fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}
