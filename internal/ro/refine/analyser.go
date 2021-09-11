package refine

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WideCount struct {
	s string
	c int
}

type Analyser struct {
}

func NewAnalyse() *Analyser {
	return &Analyser{}
}

func (a *Analyser) analyse(file string, wide, startLevel int, start *Resource) error {
	if readFile, err := os.Open(file); err != nil {
		return err
	} else {
		defer readFile.Close()

		fmt.Printf("开始 - 精炼等级: %d, %s\n", startLevel, start)

		level, times, cost := startLevel, 0, &Resource{}
		succeed, failed, broken := make([]int, 15), make([]int, 15), make([]int, 15)
		var modes []string

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			line := fileScanner.Text()
			parts := strings.Split(line, ",")

			if len(parts) == 2 {
				mode := parts[1]
				if currentLevel, err := strconv.Atoi(parts[0]); err != nil {
					log.Fatalln(err)
				} else if level != currentLevel {
					log.Fatalf("level:%d, currentLevel: %d\n", level, currentLevel)
				} else {
					times++
					fmt.Printf("%d - 等级：%d, %s", times, level, start.sub(cost))
					cost.refine(level)
					switch mode {
					case "+":
						succeed[level]++
						level++
						modes = append(modes, "O")
					case "-":
						failed[level]++
						level--
						modes = append(modes, "X")
					case "x":
						broken[level]++
						level--
						modes = append(modes, "X")
					}
					fmt.Printf(" ---> 结果：%s, 等级：%d, %s\n", mode, level, start.sub(cost))
					if mode == "x" {
						fmt.Printf("\t修复破损 - 等级：%d, Zeny：%d, 炉灰：%d", level, start.zeny-cost.zeny, start.fire-cost.fire)
						cost.repair()
						fmt.Printf(" ---> 炉灰：%d\n", start.fire-cost.fire)
					}
					if level == 3 {
						fmt.Printf("\t安全精炼 - 等级：%d, Zeny：%d, 铝：%d", level, start.zeny-cost.zeny, start.al-cost.al)
						cost.refine(level)
						level++
						fmt.Printf(" ---> 等级：%d, Zeny：%d, 铝：%d\n", level, start.zeny-cost.zeny, start.al-cost.al)
					}
				}
			}
		}

		fmt.Printf("结束 - 精炼等级: %d, %s\n", level, cost)

		for i := 4; i < 15; i++ {
			s, f, b := succeed[i], failed[i], broken[i]
			t := s + f + b
			if t > 0 {
				fmt.Printf("精炼: %d, 次数：%d, 成功：%d, 降级：%d, 破损：%d, 成功率：%d, 降级率：%d, 破损率：%d\n", i, t, s, f, b, s*100/t, f*100/(f+b), b*100/(f+b))
			}
		}

		modeStr := strings.Join(modes, "")
		size, wideCounts := len(modeStr), make(map[string]int)
		for i := 0; i <= size-wide; i++ {
			s := modeStr[i : i+wide]
			if c, ok := wideCounts[s]; ok {
				wideCounts[s] = c + 1
			} else {
				wideCounts[s] = 1
			}
		}
		var wc []*WideCount
		for s, c := range wideCounts {
			//sl := len(s)
			//if s[sl-2:sl-1] == "O" {
			wc = append(wc, &WideCount{
				s: s,
				c: c,
			})
			//}
		}
		sort.Slice(wc, func(i, j int) bool {
			wci, wcj := *wc[i], *wc[j]
			if wci.s[:wide-1] == wcj.s[:wide-1] {
				return wci.c > wcj.c
			} else {
				return wci.s[:wide-1] > wcj.s[:wide-1]
			}
		})
		for _, item := range wc {
			fmt.Printf("%s:\t%d\t%0.2f%%\n", item.s, item.c, float32(item.c*100)/float32(times-wide+1))
		}
		return nil
	}
}

func (a *Analyser) play(file string, startLevel int, start *Resource) ([]Record, *Resource, error) {
	if readFile, err := os.Open(file); err != nil {
		return nil, nil, err
	} else {
		defer readFile.Close()

		fmt.Printf("开始 - 精炼等级: %d, %s\n", startLevel, start)

		level, times, cost := startLevel, 0, &Resource{}
		var records []Record

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			if parts := strings.Split(fileScanner.Text(), ","); len(parts) == 2 {
				result := Result(parts[1])
				if currentLevel, err := strconv.Atoi(parts[0]); err != nil {
					return nil, nil, err
				} else if level != currentLevel {
					return nil, nil, errors.Errorf("level:%d, currentLevel: %d\n", level, currentLevel)
				} else {
					fmt.Printf("%d - 等级：%d, %s", times, level, start.sub(cost))

					cost.refine(level)
					times, records = times+1, append(records, Record{level: level, result: result})
					level += result.increment()

					fmt.Printf(" ---> 结果：%s, 等级：%d, %s\n", result, level, start.sub(cost))
					if result == Broken {
						fmt.Printf("\t修复破损 - 等级：%d, Zeny：%d, 炉灰：%d", level, start.zeny-cost.zeny, start.fire-cost.fire)
						cost.repair()
						fmt.Printf(" ---> 炉灰：%d\n", start.fire-cost.fire)
					}
					if level == 3 {
						fmt.Printf("\t安全精炼 - 等级：%d, Zeny：%d, 铝：%d", level, start.zeny-cost.zeny, start.al-cost.al)
						cost.refine(level)
						level++
						fmt.Printf(" ---> 等级：%d, Zeny：%d, 铝：%d\n", level, start.zeny-cost.zeny, start.al-cost.al)
					}
				}
			}
		}

		fmt.Printf("结束 - 精炼等级: %d, 花费: %s\n", level, cost)

		return records, cost, nil
	}
}
