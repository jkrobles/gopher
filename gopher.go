package gopher
/*
Title: gopher
Authors: Justice Robles
Email: justice@justicerobles.com
Website: www.justicerobles.com
Description: Gopher contains a series of useful sabermetrics functions
*/

// Slugging Percentage

/*
* Function: Slg
* Description: Calculate Slug statistic
* @param float32 TB
* @param float32 AB
* @return float32 slugging 
*/
func Slg(TB float32, AB float32 ) (slugging float32) {
  slugging = (TB/AB)
  return slugging
}

/*
* Function: Odp
* Description: onbase percentage (Odp)
* @param float32 H
* @param float32 BB
* @param float32 HBP
* @param float32 AB
* @param float32 SF
* @return float32 onbase
*/
func Odp(H float32, BB float32, HBP float32, AB float32, SF float32) (onbase float32) {
    onbase = ((H+BB+HBP)/(AB+BB+SF+HBP))
    return
}

