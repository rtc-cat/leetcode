/**
 * @file heap.cpp   
 * @author caobo.krizsx (caobo.krizsx@bytedance.com)
 * @brief 堆的实现
 * @version 0.1
 * @date 2020-11-30
 * 
 * @copyright Copyright (c) 2020
 * 
 */
#include "heap.h"

#include <iostream>
#include <vector>

// template <typename T>
// std::ostream& operator<<(std::ostream& os, const Heap<T>& heap) {
//     for (T x : heap.data_) {
//         os << x << ",";
//     }
//     os << "\n";
//     return os;
// }
#if 0
int main(int argc, char **argv)
{
    std::vector<int> vec = {55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9};
    Heap<int> heap(vec);
    std::cout << heap << std::endl;
    size_t s = heap.size();
    for (int i = 0; i < s; i++)
    {
        std::cout << heap.pop() << std::endl;
    }
}
#endif