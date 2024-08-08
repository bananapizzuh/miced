package main

import (
	"errors"

	"github.com/bananapizzuh/miced/pkg/modrinth"
)

type Platforms int

const (
	Modrinth = iota
	CurseForge
	TechnicPack
)

var PlatformNames = map[Platforms]string{
	Modrinth:    "modrinth",
	CurseForge:  "curseforge",
	TechnicPack: "tecnicpack",
}

func (ss Platforms) String() string {
	return PlatformNames[ss]
}

func main() {
	json, err := GetModpack("cobblemon-fabric", Modrinth)
	if err != nil {
		panic(err)
	}
	println(json.Title)
}

func GetModpack(slug string, platform Platforms) (modrinth.ModrinthProject, error) {
	switch platform {
	case Modrinth:
		return modrinth.GetModrinthProject(slug), nil
	case CurseForge:
		return GetCurseForgeProject(slug), nil
	case TechnicPack:
		return modrinth.ModrinthProject{}, errors.New("TechnicPack not implemented")
	default:
		return modrinth.ModrinthProject{}, errors.New("Invalid platform")
	}
}

func GetCurseForgeProject(slug string) ModrinthProject {
	return ModrinthProject{}
}
