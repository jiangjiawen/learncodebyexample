#include <iostream>

void testPtr(int* pt){
    int b=12;
    *pt=b;
    std::cout <<"2:"<< &(*pt) << std::endl;
}

void testReference(int &p){
    int a=13;
    p=a;
}

int main(){
    int a=10;
    int *p=&a;
    
    std::cout <<"0:"<< &a << std::endl;
    std::cout <<"1:"<< &p << std::endl;
    testPtr(p);
    std::cout <<"3:"<< &p << std::endl;
    std::cout <<"4:"<< *p << std::endl;
    std::cout <<"5:"<< &a << std::endl;
    
    testReference(*p);
    std::cout <<"7:"<< *p << std::endl;
    return 0;
}