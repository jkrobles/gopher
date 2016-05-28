package gopher
/*
Title: gopher
Authors: Justice Robles
Email: justice@justicerobles.com
Website: www.justicerobles.com
Description: Gopher contains a series of useful sabermetrics functions
*/

// Slugging Percentage

func Slg(TB float32, AB float32 ) (slugging float32) {
  slugging = (TB/AB)
  return slugging
}

func Odp(H float32, BB float32, HBP float32, AB float32, SF float32) (onbase float32) {
    onbase = ((H+BB+HBP)/(AB+BB+SF+HBP))
    return
}

