#include "vehicle_purchase.h"

namespace vehicle_purchase {

// needs_license determines whether a license is needed to drive a type of
// vehicle. Only "car" and "truck" require a license.
bool needs_license(std::string kind) {
    if (kind == "car" || kind == "truck") {
        return true;
    }
    return false;
}

// choose_vehicle recommends a vehicle for selection. It always recommends the
// vehicle that comes first in lexicographical order.
std::string choose_vehicle(std::string option1, std::string option2) {
    std::string choosen = "";
    if (option1 < option2) {
        choosen = option1;
    } else {
        choosen = option2;
    }
    return choosen + " is clearly the better choice.";
}

// calculate_resell_price calculates how much a vehicle can resell for at a
// certain age.
double calculate_resell_price(double original_price, double age) {
    double rate = 0.0;
    if (age < 3){
        rate = 80.0;
    } else if (age < 10){
        rate = 70.0;
    } else {
        rate = 50.0;
    }
    return (original_price * rate/100);
}

}  // namespace vehicle_purchase
