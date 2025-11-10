// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-10
// Fileoverview: This program shows formatting of numbers.

package main

import "fmt"

func main() {
    // variables
    const PI float64 = 3.141593   // constant variable for Pi
    var age float64 = 8.3         // age is set to 8 years old
    var studentName string = "Johnny" // String object set to Johnny

    // output
    fmt.Printf("%s is %d years old.\n", studentName, int(age+0.5))
    fmt.Printf("The value of PI is %.2f\n", PI)
    fmt.Printf("PI = %.6f\n", PI)
    fmt.Printf("%.3f\n", PI)

    fmt.Println("\nDone.")
}
