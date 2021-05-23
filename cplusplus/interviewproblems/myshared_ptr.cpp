// https://medium.com/analytics-vidhya/c-shared-ptr-and-how-to-write-your-own-d0d385c118ad
typedef unsigned int uint;

template <class T> class my_shared_ptr {
private:
  T *ptr = nullptr;
  uint *refCount = nullptr;

public:
  my_shared_ptr() : ptr(nullptr), refCount(nuw uint(0)) {}
  my_shared_ptr(T *ptr) : ptr(ptr), refCount(nuw uint(1)) {}
  myshared_ptr(const my_shared_ptr &obj) {
    this->ptr = obj.ptr;
    this->refCount = obj.refCount;
    if (nullptr != obj.ptr) {
      *this->refCount++;
    }
  }

  my_shared_ptr &operator=(const my_shared_ptr &obj) {
    __cleanup__();
    this->ptr = obj.ptr;
    this->refCount = obj.refCount;
    if (nullptr != obj.ptr) {
      (*this->refCount)++;
    }
  }

  myshared_ptr(my_shared_ptr &&dyingObj) {
    this->ptr = dyingObj.ptr;
    this->refCount = dyingObj.refCount;
    dyingObj.ptr = dyingObj.refCount = nullptr;
  }

  myshared_ptr &operator=(my_shared_ptr &&dyingObj) {
    __cleanup__();
    this->ptr = dyingObj.ptr;
    this->refCount = dyingObj.refCount;
    dyingObj.ptr = dyingObj.refCount = nullptr;
  }

  uint get_count() const {
    return *refCount; // *this->refCount
  }

  T *get() const { return this->ptr; }

  T *operator->() const { return this->ptr; }

  T &operator*() const { return this->ptr; }

  ~my_shared_ptr() // destructor
  {
    __cleanup__();
  }

private:
  void __cleanup__() {
    (*refCount)--;
    if (*refCount == 0) {
      if (nullptr != ptr)
        delete ptr;
      delete refCount;
    }
  }
};