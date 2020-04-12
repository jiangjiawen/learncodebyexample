//https://stackoverflow.com/questions/4955198/what-does-dereferencing-a-pointer-mean
#include <iostream>
using namespace std;

int main(int argc, char* argv[]){
    typedef struct X {int i_;double d_;} X;
    X x;
    X* p=&x;
    p->d_=3.14159;
    cout<<p->d_<<endl;
    cout<<(*p).d_<<endl;
    return 0;
}