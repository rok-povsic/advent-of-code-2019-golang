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

		printState(moons, steps)
		updateGravity(moons)
		applyVelocity(moons)
	}

	fmt.Printf("After %d steps:\n", steps)
	printState(moons, steps)

	return totalEnergy(moons)
}

func printState(moons []Moon, steps int) {
	for _, moon := range moons {
		fmt.Println(moon)
	}
}
func updateGravity(moons []Moon) {
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			if moons[i].position.x < moons[j].position.x {
				moons[i].velocity.x++
				moons[j].velocity.x--
			} else if moons[i].position.x > moons[j].position.x {
				moons[i].velocity.x--
				moons[j].velocity.x++
			}

			if moons[i].position.y < moons[j].position.y {
				moons[i].velocity.y++
				moons[j].velocity.y--
			} else if moons[i].position.y > moons[j].position.y {
				moons[i].velocity.y--
				moons[j].velocity.y++
			}

			if moons[i].position.z < moons[j].position.z {
				moons[i].velocity.z++
				moons[j].velocity.z--
			} else if moons[i].position.z > moons[j].position.z {
				moons[i].velocity.z--
				moons[j].velocity.z++
			}
		}
	}
}

func applyVelocity(moons []Moon) {
	for i := 0; i < len(moons); i++ {
		moons[i].position.x += moons[i].velocity.x
		moons[i].position.y += moons[i].velocity.y
		moons[i].position.z += moons[i].velocity.z
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

func main() {
	moons := []Moon{
		MakeMoon(17, -12, 13),
		MakeMoon(2, 1, 1),
		MakeMoon(-1, -17, 7),
		MakeMoon(12, -14, 18),
	}
	totalEnergy := Simulate(moons, 1000)
	fmt.Printf("Total energy: %d\n", totalEnergy)
}
