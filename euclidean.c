#include <stdio.h>

int euclide(int n1, int n2);

int main(int argc, char** argv)
{
    printf("%i is the gcd of 328 and 288\n",euclide(11,25));
}

int euclide(int n1, int n2)
{
    while (n1 != n2)
    {
        if (n1 > n2)
            n1 = n1 - n2;
        else
            n2 = n2 - n1;
    }

    // gcf(greatest common factor) can either be equal to n1 or n2 since they will end up with same values.

    int GCF = n1;

    return GCF;
}
