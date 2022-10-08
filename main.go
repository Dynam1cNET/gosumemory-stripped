package main

import (
	"fmt"
	"github.com/Dynam1cNET/gosumemory-stripped/db"
	"log"
	"os"
	"runtime"

	"github.com/Dynam1cNET/gosumemory-stripped/mem"
	"github.com/Dynam1cNET/gosumemory-stripped/memory"
	"github.com/Dynam1cNET/gosumemory-stripped/pp"
)

func main() {
	cgo := false
	mem.Debug = false
	memory.MemCycle = false
	memory.UpdateTime = 0
	memory.SongsFolderPath = "auto"
	memory.UnderWine = false
	if runtime.GOOS != "windows" && memory.SongsFolderPath == "auto" {
		log.Fatalln("Please specify path to osu!Songs (see --help)")
	}
	if memory.SongsFolderPath != "auto" {
		if _, err := os.Stat(memory.SongsFolderPath); os.IsNotExist(err) {
			log.Fatalln(`Specified Songs directory does not exist on the system! (try setting to "auto" if you are on Windows or make sure that the path is correct)`)
		}
	}

	go memory.Init()
	err := db.InitDB()
	if err != nil {
		log.Println(err)
		// 	time.Sleep(5 * time.Second)
		// 	os.Exit(1)
	}

	if !cgo {
		go pp.GetData()
		go pp.GetFCData()
		go pp.GetMaxData()
		go pp.GetEditorData()
	}
	//var oldLen = 0
	//var count = 0
	//for {
	//	if memory.MenuData.OsuStatus == 2 {
	//		for memory.GameplayData.KeyOverlay.K1.IsPressed || memory.GameplayData.KeyOverlay.K2.IsPressed || memory.GameplayData.KeyOverlay.M1.IsPressed || memory.GameplayData.KeyOverlay.M2.IsPressed {
	//			//log.Println(fmt.Sprint(memory.GameplayData.PP.Pp))
	//			//log.Println(fmt.Sprint(memory.MenuData.OsuStatus))
	//			if len(memory.GameplayData.Hits.HitErrorArray) != oldLen {
	//				count++
	//				oldLen = len(memory.GameplayData.Hits.HitErrorArray)
	//				if oldLen > 0 {
	//					log.Println(
	//						fmt.Sprintf("|%-6d", memory.GameplayData.Hits.HitErrorArray[oldLen-1]),
	//						fmt.Sprintf("|%-6d", len(memory.GameplayData.Hits.HitErrorArray)),
	//						fmt.Sprintf("|%-6d", count),
	//						fmt.Sprintf("|%6dpp", memory.GameplayData.PP.Pp))
	//				} else {
	//					count = 0
	//				}
	//
	//			}
	//		}
	//	} else {
	//		oldLen = 0
	//		count = 0
	//
	//	}
	//}
	for {
		if memory.MenuData.OsuStatus == 2 {
			if memory.GameplayData.KeyOverlay.K1.IsPressed || memory.GameplayData.KeyOverlay.M1.IsPressed {
				fmt.Printf("\r|%-3s|%-3s|", "###", "")
				continue
			}
			if memory.GameplayData.KeyOverlay.K2.IsPressed || memory.GameplayData.KeyOverlay.M2.IsPressed {
				fmt.Printf("\r|%-3s|%-3s|", "", "###")
				continue
			}
			if memory.GameplayData.KeyOverlay.K1.IsPressed && memory.GameplayData.KeyOverlay.K2.IsPressed || memory.GameplayData.KeyOverlay.M1.IsPressed && memory.GameplayData.KeyOverlay.M2.IsPressed {
				fmt.Printf("\r|%-3s|%-3s|", "###", "###")
				continue
			}
			if !memory.GameplayData.KeyOverlay.K1.IsPressed && !memory.GameplayData.KeyOverlay.K2.IsPressed && !memory.GameplayData.KeyOverlay.M1.IsPressed && !memory.GameplayData.KeyOverlay.M2.IsPressed {
				fmt.Printf("\r|%-3s|%-3s|", "", "")
				continue
			}

		}
	}

}
