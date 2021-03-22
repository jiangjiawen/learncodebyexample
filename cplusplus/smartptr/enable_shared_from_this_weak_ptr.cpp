#include <iostream>
#include <memory>

using namespace std;

class A:public enable_shared_from_this<A>{
    public:
    A(){
        m_i=9;
        std::cout << "A constructor" << std::endl;
    }
    ~A(){
        m_i=0;
        std::cout << "A destructor" << std::endl;
    }
    void func(){
        // m_SelfPtr = std::shared_from_this();
        m_SelfPtr = std::weak_from_this();
    }
    private:
    int m_i;
    weak_ptr<A> m_SelfPtr;
};

int main(){
    {
        shared_ptr<A> spa(new A());
        spa->func();
    }
    return 0;
}