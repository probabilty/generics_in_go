#include <iostream>
#include <string>
using namespace std;
template <typename T>

T Add(T x, T y)
{
    return x+y;
}
int main(){
    cout << Add<int>(3, 7) << endl;
    cout << Add<float>(3.0, 7.0) << endl;
    cout << Add<string>("3.0", "7.0") << endl;
    return 0;
}