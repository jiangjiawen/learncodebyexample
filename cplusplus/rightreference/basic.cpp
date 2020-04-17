// 现代c++教程：高速上手c++11/14/17/20
#include <iostream>
#include <string>

void reference(std::string &str)
{
    std::cout << "左值" << std::endl;
}

void reference(std::string &&str)
{
    std::cout << "右值" << std::endl;
}

int main()
{
    std::string lv1 = "string";
    std::string &&rv1 = std::move(lv1);
    std::cout << rv1 << std::endl;
    std::cout << rv1 << std::endl;

    const std::string &lv2 = lv1 + lv1;
    std::cout << lv2 << std::endl;

    std::string&& rv2 = lv1 + lv2;
    rv2 += "Test";
    std::cout << rv2 << std::endl;

    reference(rv2);

    return 0;
}