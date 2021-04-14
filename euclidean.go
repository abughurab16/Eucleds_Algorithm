module eucleidan

fn abs (a int) (int) {
    if a < 0 {
    	return -a
    }
    return a
}

fn eucledian_sub(n1_ int, n2_ int) (int) {
    mut n1 := abs(n1_)
    mut n2 := abs(n2_)
    
    for n1 != n2 {
        if n1 < n2 {
        	n1 = n1 - n2
        } else {
        	n2 = n2 - n1
        }
    }
}

fn eucledian_mod( n1_ int,  n2_ int) (int) {
    mut n1 := abs(n1_)
    mut n2 := abs(n2_)
    
    for n2 != 0 {
    	n2,n1 = n1%n2, n2
    }
	
    return n1
}


fn main(){
    testing := eucledian_algo(55,320)
    println(testing)
}
