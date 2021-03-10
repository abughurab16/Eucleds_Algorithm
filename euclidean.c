#include <stdio.h>

int euclide(int n1, int n2);

int main(int argc, char** argv)
{
    printf("%i\n",euclide(26,26));
}

int euclide(int n1, int n2)
{
    int GCF = 0;

    while (n1 != n2)
    {
        if (n1 > n2)
            n1 = n1 - n2;
        else
            n2 = n2 - n1;
    }
    
    // if the two numbers are equal it returns either one of the two back.  
    GCF = n1;
}
