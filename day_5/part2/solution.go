package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maps = make(map[string][][3]int)
var seeds []int

func main() {
	lines := readInput("day_5/input.txt")
	seeds = parseSeeds(lines[0])
	parseMaps(lines[1:])

	var lowestLocation int
	for _, seed := range seeds {
		location := getLocation(seed)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
		fmt.Println("seed", seed, "->", location)
	}
	fmt.Println("Lowest Location", lowestLocation)
}

func parseSeeds(line string) []int {
	re := regexp.MustCompile(`\d+`)
	hits := re.FindAllString(line, -1)
	seeds := make([]int, len(hits))
	for i, seed := range hits {
		seeds[i] = mustConvertToInt(seed)
	}

	return seeds
}

func parseMaps(lines []string) {
	mapRegexp := regexp.MustCompile(`(\S+) map:`)
	rowRegexp := regexp.MustCompile(`\d+`)
	var currentMap string

	for _, line := range lines {

		mapName := mapRegexp.FindStringSubmatch(line)
		if len(mapName) > 0 {
			currentMap = mapName[1]
			continue
		}
		mapData := rowRegexp.FindAllString(line, -1)
		if len(mapData) != 3 {
			continue
		}
		addToMap(currentMap, mustConvertToInt(mapData[0]), mustConvertToInt(mapData[1]), mustConvertToInt(mapData[2]))
	}
	//hits := mapRegexp.FindAllString(lines., -1)
	//fmt.Println(hits)

	return
}

func getLocation(seed int) int {
	value := seed

	value = getMappedValue("seed-to-soil", value)
	value = getMappedValue("soil-to-fertilizer", value)
	value = getMappedValue("fertilizer-to-water", value)
	value = getMappedValue("water-to-light", value)
	value = getMappedValue("light-to-temperature", value)
	value = getMappedValue("temperature-to-humidity", value)
	value = getMappedValue("humidity-to-location", value)

	return value
}

func addToMap(mapName string, sourceStart, destStart, length int) {
	if maps[mapName] == nil {
		maps[mapName] = make([][3]int, 0)
	}
	maps[mapName] = append(maps[mapName], [3]int{sourceStart, destStart, length})
}

func getMappedValue(mapName string, number int) int {
	for _, mapping := range maps[mapName] {
		if number >= mapping[1] && number <= mapping[1]+mapping[2] {
			return mapping[0] + number - mapping[1]
		}
	}
	return number
}

func mustConvertToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func readInput(name string) []string {
	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
