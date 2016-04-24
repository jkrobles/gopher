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
