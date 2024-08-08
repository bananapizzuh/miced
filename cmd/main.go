package main

import modrinth "miced/modrinth"

func main() {
	ModrinthClient := modrinth.NewClient("", "miced")
	project, err := ModrinthClient.GetProject("cobblemon-fabric")
	if err != nil {
		panic(err)
	}
	println(project.Title)

	projects, err := ModrinthClient.GetProjects([]string{"cobblemon-fabric", "fabric"})
	if err != nil {
		panic(err)
	}

}
