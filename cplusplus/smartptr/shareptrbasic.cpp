//现代c++教程：高速上手c++11/14/17/20
#include <iostream>
#include <memory>

void foo(std::shared_ptr<int> i)
{
    (*i)++;
}

int main()
{
    auto pointer = std::make_shared<int>(10);
    foo(pointer);
    std::cout << *pointer << std::endl;
    // The shared_ptr will be destructed before leaving the scope
    return 0;
}