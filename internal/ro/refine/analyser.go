package refine

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WideCount struct {
	s string
	c int
}

type WideCountPair struct {
	s string
	c int
	o *WideCount
	x *WideCount
}

func (p *WideCountPair) String() string {
	if p.o != nil {
		return fmt.Sprintf("%s-O: %0.2f%%, %d", p.s, float32(p.o.c*100)/float32(p.c), p.c)
	} else if p.x != nil {
		return fmt.Sprintf("%s-X: %0.2f%%, %d", p.s, float32(p.x.c*100)/float32(p.c), p.c)
	} else {
		return fmt.Sprintf("%s: %d", p.s, p.c)
	}
}

type History struct {
	file       string
	startLevel int
	start      *Resource
	times      int
	endLevel   int
	cost       *Resource
	records    []Record
}

type Analyser struct {
	histories map[string]*History
}

func NewAnalyse() *Analyser {
	return &Analyser{histories: make(map[string]*History)}
}

func (a *Analyser) analyse(wide int) error {
	times, succeed, failed, broken, wideCounts := 0, make([]int, 15), make([]int, 15), make([]int, 15), make(map[string]int)

	for _, history := range a.histories {
		times += history.times
		var modes []string
		for _, record := range history.records {
			switch record.result {
			case Succeed:
				succeed[record.level]++
				modes = append(modes, "O")
			case Failed:
				failed[record.level]++
				modes = append(modes, "X")
			case Broken:
				broken[record.level]++
				modes = append(modes, "X")
			}
		}

		size, modeStr := len(modes), strings.Join(modes, "")
		for i := 0; i <= size-wide; i++ {
			s := modeStr[i : i+wide]
			if c, ok := wideCounts[s]; ok {
				wideCounts[s] = c + 1
			} else {
				wideCounts[s] = 1
			}
		}
	}

	for i := 4; i < 15; i++ {
		s, f, b := succeed[i], failed[i], broken[i]
		t := s + f + b
		if t > 0 {
			fmt.Printf("精炼：%d, 次数：%d, 成功：%d, 降级：%d, 破损：%d, 成功率：%0.2f%%, 降级率：%0.2f%%, 破损率：%0.2f%%\n",
				i, s, t, f, b, float32(s*100)/float32(t), float32(f*100)/float32(f+b), float32(b*100)/float32(f+b))
		}
	}

	count, pairs := 0, make(map[string]*WideCountPair)
	for s, c := range wideCounts {
		count += c
		prefix, suffix := s[:wide-1], s[wide-1:]
		pair, ok := pairs[prefix]
		if !ok {
			pair = &WideCountPair{s: prefix}
			pairs[prefix] = pair
		}
		pair.c += c
		if "O" == suffix {
			pair.o = &WideCount{
				s: s,
				c: c,
			}
		} else {
			pair.x = &WideCount{
				s: s,
				c: c,
			}
		}
	}

	var pairsArray []*WideCountPair
	for _, pair := range pairs {
		pairsArray = append(pairsArray, pair)
	}
	sort.Slice(pairsArray, func(i, j int) bool {
		pairi, pairj := *pairsArray[i], *pairsArray[j]
		if pairi.o != nil && pairj.o == nil {
			return true
		} else if pairi.o == nil && pairj.o != nil {
			return false
		} else if pairi.o == nil && pairj.o == nil {
			return pairi.c > pairj.c
		} else {
			ratei := pairi.o.c * 10000 / pairi.c
			ratej := pairj.o.c * 10000 / pairj.c
			return ratei > ratej
		}
	})
	for _, pair := range pairsArray {
		fmt.Printf("%0.2f%% - %s\n", float32(pair.c*100)/float32(count), pair)
	}
	return nil
}

func (a *Analyser) play(file string, startLevel int, start *Resource, debug bool) (*History, error) {
	if readFile, err := os.Open(file); err != nil {
		return nil, err
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
					return nil, err
				} else if level != currentLevel {
					return nil, errors.Errorf("level:%d, currentLevel: %d\n", level, currentLevel)
				} else {
					if debug {
						fmt.Printf("%d - 等级：%d, %s", times, level, start.sub(cost))
					}

					cost.refine(level)
					times, records = times+1, append(records, Record{level: level, result: result})
					level += result.increment()

					if debug {
						fmt.Printf(" ---> 结果：%s, 等级：%d, %s\n", result, level, start.sub(cost))
					}
					if result == Broken {
						if debug {
							fmt.Printf("\t修复破损 - 等级：%d, Zeny：%d, 炉灰：%d", level, start.zeny-cost.zeny, start.fire-cost.fire)
						}
						cost.repair()
						if debug {
							fmt.Printf(" ---> 炉灰：%d\n", start.fire-cost.fire)
						}
					}
					if level == 3 {
						if debug {
							fmt.Printf("\t安全精炼 - 等级：%d, Zeny：%d, 铝：%d", level, start.zeny-cost.zeny, start.al-cost.al)
						}
						cost.refine(level)
						level++
						if debug {
							fmt.Printf(" ---> 等级：%d, Zeny：%d, 铝：%d\n", level, start.zeny-cost.zeny, start.al-cost.al)
						}
					}
				}
			}
		}

		fmt.Printf("结束 - 精炼次数: %d, 最终等级: %d, 花费: %s\n", times, level, cost)

		history := &History{
			file:       file,
			startLevel: startLevel,
			start:      start,
			times:      times,
			endLevel:   level,
			cost:       cost,
			records:    records,
		}
		a.histories[file] = history
		return history, nil
	}
}
