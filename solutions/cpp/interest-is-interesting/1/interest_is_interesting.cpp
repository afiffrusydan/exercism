// interest_rate returns the interest rate for the provided balance.
double interest_rate(double balance) {
    double rate{0.0};
    if (balance < 0){
        rate = 3.213;
    } else if (balance >= 5000){
        rate = 2.475;
    } else if (balance >= 1000){
        rate = 1.621;
    } else {
        rate = 0.5;
    }
    return rate;
}

// yearly_interest calculates the yearly interest for the provided balance.
double yearly_interest(double balance) {
    return balance * interest_rate(balance) / 100;
}

// annual_balance_update calculates the annual balance update, taking into
// account the interest rate.
double annual_balance_update(double balance) {    
    return balance + yearly_interest(balance);
}

// years_until_desired_balance calculates the minimum number of years required
// to reach the desired balance.
int years_until_desired_balance(double balance, double target_balance) {
    int year{0};
    double new_balance{balance};
    for (int i{1}; new_balance < target_balance; ++i){
        new_balance = annual_balance_update(new_balance);
        year = i;
    }
    return year;
}
