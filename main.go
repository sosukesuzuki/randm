package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var adjectives = [...]string{"amazing", "beautiful", "charming", "delightful", "elegant", "fabulous", "graceful", "harmonious", "innovative", "joyful", "kind", "luminous", "magnificent", "noble", "optimistic", "peaceful", "quirky", "radiant", "stunning", "tranquil", "unique", "vibrant", "wonderful", "youthful", "zealous", "adorable", "brilliant", "creative", "dazzling", "exquisite", "fantastic", "gorgeous", "hilarious", "inspiring", "jovial", "keen", "lively", "marvelous", "nifty", "outstanding", "pleasant", "quaint", "remarkable", "spectacular", "terrific", "unforgettable", "vivid", "whimsical", "youthful", "zesty", "affable", "breezy", "crisp", "daring", "effervescent", "festive", "glamorous", "heartwarming", "intuitive", "jubilant", "knowledgeable", "luxurious", "mystical", "notable", "opulent", "picturesque", "quintessential", "resplendent", "serene", "thrilling", "upbeat", "venerable", "witty", "xenial", "yielding", "zany", "artistic", "blissful", "captivating", "dynamic", "energetic", "flawless", "grateful", "hopeful", "idealistic", "jazzy", "kind-hearted", "legendary", "majestic", "nimble", "observant", "pristine", "quintessential", "rustic", "sophisticated", "timeless", "uplifting", "versatile", "wholesome", "xanadu", "yummy", "zippy"}
var names = [...]string{"naruto", "luffy", "goku", "saitama", "light", "eren", "mikasa", "sasuke", "ichigo", "natsu", "tanjiro", "nezuko", "edward", "alphonse", "astaroth", "bulma", "chihiro", "deku", "erza", "frieza", "gintoki", "hiei", "inosuke", "jojo", "kirito", "lelouch", "motoko", "nami", "oliver", "piccolo", "quinty", "rem", "shinji", "totoro", "usagi", "vegeta", "wiz", "xanxus", "yato", "zoro", "asuna", "broly", "celty", "doraemon", "elric", "fujiko", "genos", "hinata", "itachi", "jiraiya", "kaiba", "kakashi", "l", "mario", "nico", "oda", "pikachu", "quinn", "ryuk", "sakura", "takagi", "umaru", "vash", "winry", "xander", "yugi", "zenitsu", "ace", "boa", "crona", "diane", "escanor", "fubuki", "grell", "haruhi", "izuku", "jolyne", "kenpachi", "law", "mugen", "nagisa", "orochimaru", "portgas", "qrow", "riku", "senku", "taiga", "ulquiorra", "violet", "watanuki", "xeno", "yamcha", "zabuza", "akame", "ban"}

func main() {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	adj := adjectives[r.Intn(len(adjectives))]
	name := names[r.Intn(len(names))]
	fmt.Fprint(os.Stdout, adj+"-"+name)
}
