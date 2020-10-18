package main

import (
	"fmt"
)

type Coordinates struct {
	x, y, z int
}

type Moon struct {
	position, velocity Coordinates
}

func MakeMoon(x int, y int, z int) Moon {
	moon := Moon{}
	moon.position = Coordinates{x, y, z}
	moon.velocity = Coordinates{0, 0, 0}
	return moon
}

func Simulate(moons []Moon, steps int) int {
	for step := 0; step < steps; step++ {
		fmt.Printf("After %d steps:\n", step)

		stepX(moons)
		stepY(moons)
		stepZ(moons)
	}

	fmt.Printf("After %d steps:\n", steps)
	printState(moons, steps)

	return totalEnergy(moons)
}

func RepeatSteps(initialMoons []Moon, moons []Moon) int {
	stepsX := 1
	for true {
		stepX(moons)

		equal := true
		for i := 0; i < len(moons); i++ {
			if moons[i].position.x != initialMoons[i].position.x || moons[i].velocity.x != initialMoons[i].velocity.x {
				equal = false
				break
			}
		}
		if equal {
			break
		}
		stepsX++
	}

	stepsY := 1
	for true {
		stepY(moons)

		equal := true
		for i := 0; i < len(moons); i++ {
			if moons[i].position.y != initialMoons[i].position.y || moons[i].velocity.y != initialMoons[i].velocity.y {
				equal = false
				break
			}
		}
		if equal {
			break
		}
		stepsY++
	}

	stepsZ := 1
	for true {
		stepZ(moons)

		equal := true
		for i := 0; i < len(moons); i++ {
			if moons[i].position.z != initialMoons[i].position.z || moons[i].velocity.z != initialMoons[i].velocity.z {
				equal = false
				break
			}
		}
		if equal {
			break
		}
		stepsZ++
	}

	fmt.Println(stepsX)
	fmt.Println(stepsY)
	fmt.Println(stepsZ)

	fmt.Printf("After %d steps:\n", stepsX)
	steps := lcm(stepsX, stepsY, stepsZ)
	printState(moons, steps)
	return steps
}

func stepX(moons []Moon) {
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].position.x < moons[j].position.x {
				moons[i].velocity.x++
				moons[j].velocity.x--
			} else if moons[i].position.x > moons[j].position.x {
				moons[i].velocity.x--
				moons[j].velocity.x++
			}
		}
	}

	for i := 0; i < len(moons); i++ {
		moons[i].position.x += moons[i].velocity.x
	}
}

func stepY(moons []Moon) {
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].position.y < moons[j].position.y {
				moons[i].velocity.y++
				moons[j].velocity.y--
			} else if moons[i].position.y > moons[j].position.y {
				moons[i].velocity.y--
				moons[j].velocity.y++
			}
		}
	}

	for i := 0; i < len(moons); i++ {
		moons[i].position.y += moons[i].velocity.y
	}
}

func stepZ(moons []Moon) {
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].position.z < moons[j].position.z {
				moons[i].velocity.z++
				moons[j].velocity.z--
			} else if moons[i].position.z > moons[j].position.z {
				moons[i].velocity.z--
				moons[j].velocity.z++
			}
		}
	}

	for i := 0; i < len(moons); i++ {
		moons[i].position.z += moons[i].velocity.z
	}
}

func printState(moons []Moon, steps int) {
	for _, moon := range moons {
		fmt.Println(moon)
	}
}
func totalEnergy(moons []Moon) int {
	result := 0
	for _, moon := range moons {
		potentialEnergy := abs(moon.position.x) + abs(moon.position.y) + abs(moon.position.z)
		kineticEnergy := abs(moon.velocity.x) + abs(moon.velocity.y) + abs(moon.velocity.z)
		result += potentialEnergy * kineticEnergy
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func main() {
	initialMoons := []Moon{
		MakeMoon(17, -12, 13),
		MakeMoon(2, 1, 1),
		MakeMoon(-1, -17, 7),
		MakeMoon(12, -14, 18),
	}

	moons := []Moon{
		MakeMoon(17, -12, 13),
		MakeMoon(2, 1, 1),
		MakeMoon(-1, -17, 7),
		MakeMoon(12, -14, 18),
	}
	//totalEnergy := Simulate(moons, 1000)
	//fmt.Printf("Total energy: %d\n", totalEnergy)
	//
	//fmt.Println(initialMoons)
	//fmt.Println(moons)

	steps := RepeatSteps(initialMoons, moons)
	fmt.Printf("Number of steps: %d\n", steps)
}
