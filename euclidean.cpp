#include <iostream>

void euclide(int n1, int n2);

int main(int argc, char** argv)
{
   euclide(20,10);
}

void euclide(int n1, int n2)
{
    int GCF = 0;
    int iterator = 0;
    
    while (n1 != n2)
    {
        if (n1 > n2)
            n1 = n1 - n2;
        else
            n2 = n2 - n1;
        
        iterator = iterator + 1;
    }
    
    // if the two numbers are equal it returns either one of the two back.  
    GCF = n1;
    
    std::cout << "ran " << iterator << " times\n";
    std::cout << "GCF ==> " << GCF << "\n";
}
