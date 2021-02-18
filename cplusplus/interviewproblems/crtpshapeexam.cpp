//https://zhuanlan.zhihu.com/p/348075456
#include <iostream>

template <typename T> class Shape {
private:
  struct accessor : public T {
    double __get_area() { return this->T::get_area_impl(); }
  };

public:
  double get_area() { return static_cast<accessor &>(*this)->__get_area(); }
};

class Circle : public Shape<Circle> {
private:
  double radius = 1;

protected:
  double get_area_impl() { return 3.1415926 * radius * radius; }
};

class Rectangle : public Shape<Rectangle> {
private:
  double a = 1, b = 2;

protected:
  double get_area_impl() { return a * b; };
};

template <typename Sp> double get_area(Shape<Sp> *sp) { return sp->get_area(); }

int main() {
  Shape<Circle> *c = new Circle;
  Shape<Rectangle> *r = new Rectangle;
  std::cout << get_area(c) << " " << get_area(r) << std::endl;
  return 0;
}