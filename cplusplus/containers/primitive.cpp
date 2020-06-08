#include <iostream>
#include <vector>
using namespace std;

int main(){
    vector<int> v{1,2,3,4,5,6,7,8};
    vector<int> *v2 = &v;
    // vector<int> v2 = v;
    v2->push_back(9);
    cout << v.at(8) << endl;
    cout << v[8] << endl;
    return 0;
}