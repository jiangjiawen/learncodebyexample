#include <iostream>

void reference(int& v){
    std::cout << "left value" <<std::endl;
}

void reference(int&& v){
    std::cout << "right value" <<std::endl;
}

template <typename T>
void pass(T&& v){
    std::cout << "common pass parameters" <<std::endl;
    reference(v);
}

int main(){
    std::cout << "pass right value" <<std::endl;
    pass(1);

    std::cout << "pass left value" <<std::endl;
    int l=1;
    pass(l);

    return 0;
}