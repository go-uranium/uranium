package controllor

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Mod map[uint]bool // nid list

func NewMod(str string) (*Mod,error) {
	mod := make(Mod)
	str = strings.TrimSpace(str)
	ModList := strings.Split(str, "|")
	for _,v := range ModList {
		nid,err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		mod[uint(nid)] = true
	}
	return &mod,nil
}

func (mod *Mod) MarshalJSON() ([]byte,error) {
	var ml []uint
	for k :=range *mod {
		ml = append(ml,k)
	}
	return json.Marshal(ml)
}