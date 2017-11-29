package helpers

import "github.com/bwmarrin/discordgo"

func RemoveDuplicateUsers(users []*discordgo.User) []*discordgo.User{
	seen := make(map[string]discordgo.User)
	j := 0
	for _, u := range users{
		if _, ok := seen[u.ID]; ok{
			continue
		}
		seen[u.ID] = *u
		users[j] = u
		j++
	}
	return users[:j]
}

func GiveUserAndDiscrim(u *discordgo.User) string{
	return u.Username+"#"+u.Discriminator
}
