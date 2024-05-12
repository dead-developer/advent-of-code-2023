package main

import (
	"AoC2023/framework"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

var maps = make(map[string][][3]int)
var seeds [][2]int

func main() {
	value := solution()
	fmt.Println("Lowest Location", value)
}

func solution() int {
	lines := framework.ReadInput("input.txt")

	seeds = parseSeeds(lines[0])
	parseMaps(lines[1:])

	var lowestLocation int

	locationCh := make(chan int, len(seeds))
	for _, seedRange := range seeds {
		go getLowest(seedRange, locationCh)
	}

	for range seeds {
		location := 0
		location = <-locationCh
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func getLowest(seedRange [2]int, lowestLocations chan int) {
	var lowestLocation int
	for seed := seedRange[0]; seed <= seedRange[1]; seed++ {
		location := getLocation(seed)
		if lowestLocation == 0 || location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println("Lowest location for range:", seedRange[0], "-", seedRange[0]+seedRange[1], "is", lowestLocation)
	lowestLocations <- lowestLocation
}

func parseSeeds(line string) [][2]int {
	re := regexp.MustCompile(`\d+`)
	hits := re.FindAllString(line, -1)

	// seed ranges
	for i := 0; i < len(hits); i += 2 {
		start := mustConvertToInt(hits[i])
		end := mustConvertToInt(hits[i+1]) + start
		seedRange := [2]int{start, end}
		seeds = append(seeds, seedRange)
	}
	//sort seeds by start
	sort.Slice(seeds, func(i, j int) bool {
		return seeds[i][0] < seeds[j][0]
	})

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
