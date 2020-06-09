// https://www.youtube.com/watch?v=_8-ht2AKyH4
#include <stdio.h>
#include <stdlib.h>

int main() {
  int a;
  int *p;
  p = new int;
  *p = 10;
  delete p;
  p = new int[20];
  delete[] p;
  return 0;
}