//https://www.geeksforgeeks.org/write-memcpy/
//https://www.tutorialspoint.com/write-your-own-memcpy-and-memmove-in-cplusplus
#include <stdio.h>
#include <string>

using namespace std;

void myMemCpy(void *dest,void *src, size_t n)
{
    char *csrc=(char *)src;
    char *cdest = (char *)dest;
    for(int i=0;i<n;i++){
        cdest[i] = csrc[i];
    }
}

void MemcpyFunc(void *dest, void *src, size_t n){
//    char *dataS = (char *)src;
//    char *dataD = (char *)dest;
   for (int i=0; i<n; i++)
      *((char *)(dest)+i) = *((char *)(src)+i);
}

void myMemMove(void *dest, void *src, size_t n)
{
   // Typecast src and dest addresses to (char *)
   char *csrc = (char *)src;
   char *cdest = (char *)dest;
  
   // Create a temporary array to hold data of src
   char *temp = new char[n];
  
   // Copy data from csrc[] to temp[]
   for (int i=0; i<n; i++)
       temp[i] = csrc[i];
  
   // Copy data from temp[] to cdest[]
   for (int i=0; i<n; i++)
       cdest[i] = temp[i];
  
   delete [] temp;
}

bool isNumericChar(char x) {
   return (x >= '0' && x <= '9') ? true : false;
}
int myAtoi(char* str) {
   if (*str == '\0')
      return 0;
   int result = 0;
   int sign = 1;
   int i = 0;
   if (str[0] == '-') {
      sign = -1;
      i++;
   }
   for (; str[i] != '\0'; ++i) {
      if (isNumericChar(str[i]) == false)
         return 0;
      result = result * 10 + str[i] - '0';
   }
   return sign * result;
}

 void my_func_strcpy(char *source, char* destination, int n)
 {
    char temp[n] = {'\0'};
    int  index    = 0;

    /* Copying the destination data to source data
     */
    while (destination[index] != '\0')
    {
        source[index] = destination[index];
        index++;
    }

    /* Making the rest of the characters null ('\0') 
     */
    for (index = 0; index < n; index++)
    {
        source[index] = '\0';
    }
 }

int main() {
   char string[] = "-32491841";
   int intVal = myAtoi(string);
   cout<<"The integer equivalent of the given string is "<<intVal;
   return 0;
}