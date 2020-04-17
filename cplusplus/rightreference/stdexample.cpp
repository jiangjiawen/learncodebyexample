#include <iostream>
#include <utility>
#include <vector>
#include <string>

int main()
{
    std::string str = "Hello world";
    std::vector<std::string> v;

    v.push_back(str);
    std::cout << "str: " << str << std::endl;

    v.push_back(std::move(str));
    std::cout << "str: " << str << std::endl;
    std::cout << v[0] << std::endl;

    return 0;
}