#pragma once
#include <iostream>
#include <vector>

template <typename T>
class Heap
{

public:
    std::vector<T> data_;
    // constructor
    Heap();
    Heap(int length);
    Heap(std::vector<T> &vec);
    Heap(const Heap &other);
    virtual ~Heap();
    // method
    void up(int index);
    void down(int index);
    size_t size();
    void push(T &x);
    T &pop();
    T &peek();
    friend std::ostream &operator<<(std::ostream &os, const Heap<T> &heap)
    {
        for (T x : heap.data_)
        {
            os << x << ",";
        }
        os << "\n";
        return os;
    }
};

template <typename T>
Heap<T>::Heap()
{
}

template <typename T>
Heap<T>::Heap(int length) : data_(std::vector<T>(length))
{
}

template <typename T>
Heap<T>::Heap(std::vector<T> &vec)
{
    data_ = vec;
    size_t s = data_.size();
    // 从第一个拥有子节点的节点
    for (int i = (s / 2) - 1; i >= 0; i--)
    {
        down(i);
    }
}

template <typename T>
Heap<T>::Heap(const Heap<T> &other)
{
    data_ = other.data_;
    size_t s = data_.size();
    for (int i = (s / 2) - 1; i >= 0; i--)
    {
        down(i);
    }
}

template <typename T>
Heap<T>::~Heap()
{
    // std::cout << "析构函数" << std::endl;
}
template <typename T>
size_t Heap<T>::size()
{
    return data_.size();
}

template <typename T>
void Heap<T>::up(int index)
{
    for (;;)
    {
        // 如果节点存在
        int r = index - 1;
        if (r < 0)
        {
            break;
        }
        int parent = r / 2;
        if (data_[parent] > data_[index])
        {
            break;
        }
        else
        {
            std::swap<T>(data_[parent], data_[index]);
        }
        index = parent;
    }
}

template <typename T>
void Heap<T>::down(int index)
{
    for (;;)
    {
        int max = index;
        // 左子节点存在
        int left = index * 2 + 1;
        if (left < data_.size())
        {
            if (data_[left] > data_[max])
            {
                max = left;
            }
        }
        // 右子节点存在
        int right = index * 2 + 2;
        if (right < data_.size())
        {
            if (data_[right] > data_[max])
            {
                max = right;
            }
        }
        // 如果相等说明堆化完成
        if (max == index)
        {
            break;
        }
        else
        {
            // 否则就交换两个节点的值
            std::swap<T>(data_[max], data_[index]);
        }
        index = max;
    }
}

/**
 * @brief 将新插入的元素放在最后,并且上浮
 * 
 * @param x 
 */

template <typename T>
void Heap<T>::push(T &x)
{
    data_.push_back(x);
    up(data_.size() - 1);
}

/**
 * @brief 取出堆顶元素
 * 
 * @return int 
 */
template <typename T>
T &Heap<T>::pop()
{
    int result = data_[0];
    std::swap<T>(data_[0], data_[data_.size() - 1]); //交换第一个和最后一个元素
    data_.pop_back();                                // 删掉最后一个
    down(0);                                         // 重新堆化
    return result;
}

template <typename T>
T &Heap<T>::peek()
{
    return data_[0];
}
