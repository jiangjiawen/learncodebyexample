#include <iostream>
#include <utility>

void reference(int &v)
{
    std::cout << "left value reference" << std::endl;
}

void reference(int &&v)
{
    std::cout << "right value reference" << std::endl;
}

template <typename T>
void pass(T &&v)
{
    std::cout << "          common pass parameters:";
    reference(v);
    std::cout << "       std::move pass parameters:";
    reference(std::move(v));
    std::cout << "    std::forward pass parameters:";
    reference(std::forward<T>(v));
    std::cout << "static_cast<T&&> pass parameters:";
    reference(static_cast<T>(v));
}

int main(){
    std::cout << "pass right value" <<std::endl;
    pass(1);

    std::cout << "pass left value" <<std::endl;
    int l=1;
    pass(l);

    return 0;
}