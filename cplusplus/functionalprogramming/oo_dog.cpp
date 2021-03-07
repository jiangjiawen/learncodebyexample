#include <iostream>

class dog
{
    public:
    dog();
    bool eat(int amount_of_food);
    private:
    int hunger_;
    bool needs_shower_;
}

dog::doge():hunger_(20), needs_shower_(false) {}

bool dog::eat(int amount_of_food)
{
    if(amount_of_food <= hunger_){
        std::cout << "*tail wagging*" << std::endl;
        hunger_-=amount_of_food;
        return true;
    }else {
        std::cout << "*vomit*" << std::endl;
        needs_shower_ = true;
        return false;
    }
}

int main(){
    doge rokey;
    rokey.eat(15);
    rokey.eat(15);
}