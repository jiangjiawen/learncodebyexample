template <typename CRTP> struct SomeTemplate {
  bool operator==(const SomeTemplate &rhs) {
    return !(static_cast<const CRTP &>(*this) <
                 static_cast<const CRTP &>(rhs) ||
             static_cast<const CRTP &>(rhs) < static_cast<const CRTP &>(*this));
  }
  bool operator!=(const SomeTemplate &rhs) { return !(*this == rhs); }
};

struct MyStruct : SomeTemplate<MyStruct> {
  MyStruct(const int data_) : data(data_) {}
  int data;
  bool operator<(const MyStruct &rhs) const { return data < rhs.data; }
};

MyStruct getValue() {
  static int cur_value = 0;
  return MyStruct{++cur_value};
}

int main() { return getValue() == getValue(); }