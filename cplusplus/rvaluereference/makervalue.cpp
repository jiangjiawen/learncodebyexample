#include <iostream>
#include <string>
using namespace std;

int main(int argc, char *argv[]){
    // int x = 666;
    // int y = x + 5;
    // string s1 = "hello";
    // string s2 = "world";
    // string s3 = s1 + s2;

    // string getString(){
    //     return "hello world";
    // }
    // string s4 = getString();

    string s1 = "hello";
    string s2 = "world";
    string&& s_rref = s1 + s2;
    s_rref += ",my friend";
    cout << s_rref << endl;
    
    return 0;
}