// xs -> ys == [0,1,2,3,4] -> [0,1,4,9,16]
// https://github.com/Dobiasd/FunctionalPlus
// https://www.youtube.com/watch?v=27tFyKc6g3U&ab_channel=TobiasHermann
// compile cmd: clang++ -O3 -std=c++14 square_values.cpp -o square_values
// #include <fplus/fplus.hpp>
#include <algorithm>
#include <iterator>
#include <vector>

using namespace std;
typedef vector<int> Ints;
int square(int x) { return x * x; }

Ints square_vec_goto(const Ints &xs) {
  Ints ys;
  ys.reserve(xs.size());
  auto it = begin(xs);
loopBegin:
  if (it == end(xs)) {
    goto loopEnd;
  }
  ys.push_back(square(*it));
  ++it;
  goto loopBegin;
loopEnd:
  return ys;
}

Ints square_vec_whie(const Ints& xs)
{
    Ints ys;
    ys.reserve(xs.size());
    auto it = begin(xs);
    while(it!=end(xs)) {
        ys.push_back(square(*it));
        ++it;
    }
    return ys;
}

Ints square_vec_for(const Ints& xs)
{
    Ints ys;
    ys.reserve(xs.size());
    for(auto it=begin(xs);it!=end(xs);it++){
        ys.push_back(square(*it));
    }
    return ys;
}

Ints range_based_for(const Ints& xs)
{
    Ints ys;
    ys.reserve(xs.size());
    for(int x:xs){
        ys.push_back(square(x));
    }
    return ys;
}

Ints sqr_std_transform(const Ints& xs)
{
    Ints ys;
    ys.reserve(xs.size());
    transform(begin(xs), end(xs),back_inserter(ys),square);
    return ys;
}
template <typename F, typename T>
vector<T> transform_vec(F f, const vector<T>& xs)
{
    vector<T> ys;
    ys.reserve(xs.size());
    transform(begin(xs),end(xs),back_inserter(ys),f);
    return ys;
}
Ints sqr_transform_vec(const Ints& xs)
{
    return transform_vec(square, xs);
}

// Ints sqr_fplus_transform(const Ints& xs)
// {
//     return fplus::transform(square, xs);
// }

int main()
{
    Ints xs(8192);
    for(int i=0;i<65536;++i){
        sqr_transform_vec(xs);
    }
}