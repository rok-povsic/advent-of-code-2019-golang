package main

import (
	"testing"
)

func TestMaps(t *testing.T) {
	asteroidMap := ".#..#\n.....\n#####\n....#\n...##"
	x, y, num := BestAsteroidLocation(asteroidMap)
	if x != 3 || y != 4 || num != 8 {
		t.Errorf("Unexpected location x=%d y=%d", x, y)
	}

	asteroidMap = "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####"
	x, y, num = BestAsteroidLocation(asteroidMap)
	if x != 5 || y != 8 || num != 33 {
		t.Errorf("Unexpected location x=%d y=%d", x, y)
	}

	asteroidMap = "#.#...#.#.\n.###....#.\n.#....#...\n##.#.#.#.#\n....#.#.#.\n.##..###.#\n..#...##..\n..##....##\n......#...\n.####.###."
	x, y, num = BestAsteroidLocation(asteroidMap)
	if x != 1 || y != 2 || num != 35 {
		t.Errorf("Unexpected location x=%d y=%d", x, y)
	}

	asteroidMap = ".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#.."
	x, y, num = BestAsteroidLocation(asteroidMap)
	if x != 6 || y != 3 || num != 41 {
		t.Errorf("Unexpected location x=%d y=%d", x, y)
	}

	asteroidMap = ".#..##.###...#######\n##.############..##.\n.#.######.########.#\n.###.#######.####.#.\n#####.##.#.##.###.##\n..#####..#.#########\n####################\n#.####....###.#.#.##\n##.#################\n#####.##.###..####..\n..######..##.#######\n####.##.####...##..#\n.#####..#.######.###\n##...#.##########...\n#.##########.#######\n.####.#.###.###.#.##\n....##.##.###..#####\n.#.#.###########.###\n#.#.#.#####.####.###\n###.##.####.##.#..##"
	x, y, num = BestAsteroidLocation(asteroidMap)
	if x != 11 || y != 13 || num != 210 {
		t.Errorf("Unexpected location x=%d y=%d", x, y)
	}
}
