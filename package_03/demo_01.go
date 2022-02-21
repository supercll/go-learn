package main

import "fmt"

type Anime struct {
	title    string
	producer string
	anime_id int
}

func main() {
	var anime1 Anime /* 声明 anime1 为 Anime 类型 */
	var anime2 Anime /* 声明 anime2 为 Anime 类型 */

	/* anime 1 描述 */
	anime1.title = "咒术回战"
	anime1.producer = "MAPPA"
	anime1.anime_id = 6495407

	/* anime 2 描述 */
	anime2.title = "来自深渊"
	anime2.producer = "KINEMA CITRUS"
	anime2.anime_id = 6495700

	changeAnime(anime2)
	changeAnimeByPointer(&anime1)
	fmt.Println(anime1)
	fmt.Println(anime2)
}

func changeAnime(anime Anime) {
	anime.title = "进击的巨人"
}
func changeAnimeByPointer(anime *Anime) {
	anime.title = "黑之契约者"
}
