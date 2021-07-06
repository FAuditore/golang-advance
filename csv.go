package main

import "fmt"

func main() {
	//file, err := os.OpenFile("lzhsb.csv", os.O_RDWR|os.O_CREATE, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//file.WriteString("\xEF\xBB\xBF")

	//w := csv.NewWriter(file)
	//w.Write([]string{"path", "total_conversions", "total_value", "total_null"})

	path := []string{"A", "B", "C", "D", "E", "F", ""}

	var results [][]string

	permute(path, 0, len(path), &results)

	//fmt.Println(results)

	paths := buildPaths(results)

	fmt.Println(paths)

	//w.Flush()
}

func buildPaths(results [][]string) []string {
	var paths []string
	set := make(map[string]struct{})

	for _, v := range results {
		str := ""
		for i, c := range v {
			if i == 0 && c != "" {
				str = c
				continue
			}
			if c != "" {
				str += ">" + c
			} else {
				break
			}
		}
		set[str] = struct{}{}
	}

	for k, _ := range set {
		paths = append(paths, k)
	}
	return paths
}

func permute(path []string, start int, end int, results *[][]string) {
	if start == len(path)-1 {
		result := make([]string, len(path))
		copy(result, path)
		*results = append(*results, result)
		return
	}
	for i := start; i < end; i++ {
		path[start], path[i] = path[i], path[start]
		permute(path, start+1, end, results)
		path[start], path[i] = path[i], path[start]
	}
}
