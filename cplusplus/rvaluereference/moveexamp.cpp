// https://www.internalpointers.com/post/c-rvalue-references-and-move-semantics-beginners
#include <iostream>
using namespace std;

class Holder
{
    public:
    Holder(int size)
    {
        m_data = new int[size];
        m_size = size;
    }
    Holder(const Holder& other)
    {
        m_data = new int[other.m_size];  // (1)
        std::copy(other.m_data, other.m_data + other.m_size, m_data);  // (2)
        m_size = other.m_size;
    }

    Holder& operator=(const Holder& other) 
    {
        if(this == &other) return *this;  // (1)
        delete[] m_data;  // (2)
        m_data = new int[other.m_size];
        std::copy(other.m_data, other.m_data + other.m_size, m_data);
        m_size = other.m_size;
        return *this;  // (3)
    }

    ~Holder()
    {
        delete[] m_data;
    }
    private:
    int* m_data;
    size_t m_size;
};

int main(int argc, char* argv[]){
    Holder h1(10000);  // regular constructor
    Holder h2(60000);  // regular constructor
    Holder h3(h1);
    h1 = h2;           // assignment operator
        // Holder h3(h1);    
    return 0;
}