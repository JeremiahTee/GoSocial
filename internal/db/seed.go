package db

import (
	"context"
	"fmt"
	"github.com/JeremiahTee/GoSocial/internal/store"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var usernames = []string{
	"Rhaenyra", "Alicent", "Daemon", "Viserys", "AegonII", "Helaena", "Jacaerys", "Lucerys",
	"Rhaenys", "Corlys", "Laenor", "Laena", "Otto", "CristonCole", "Mysaria", "LarysStrong",
	"HarwinStrong", "JoffreyVelaryon", "Baela", "Rhaena", "Aemond", "Vhagar", "Sunfyre", "Caraxes",
	"Syrax", "Meleys", "Dreamfyre", "Balerion", "Silverwing", "Vermithor", "MaegorTheCruel",
	"Egg", "DunkTheTall", "DarkSister", "Bloodraven", "TheMadKing", "YoungGriff", "JonConnington",
	"Blackfyre", "DaemonBlackfyre", "QuentynMartell", "TheRedViper", "OberynMartell",
	"NymeriaSand", "ArianneMartell", "DoranMartell", "AreoHotah", "HarrenTheBlack",
	"TorrhenStark", "LyarraStark", "RickardStark", "BrandonStark", "HowlandReed",
}

var titles = []string{
	"Fire Was My Birthright…",
	"We Were Never Meant to Grow Old",
	"The Price of Loyalty Is Blood",
	"A Throne Drenched in Ash",
	"The Rats Always Knew…",
	"A Debt Paid in Fire and Blood",
	"Dreams Did Not Save Me…",
	"Whispers Weave the Fate of Kings",
	"Honor Is a Man’s Greatest Lie",
	"What Is Wealth Without Legacy?",
	"Names Written in Red",
	"A Crown Was Never Mine to Hold",
	"Power Is Power, Until It Isn’t…",
	"Wit Is a Blade Sharper Than Valyrian Steel",
	"The Cold Truth of Honor",
	"Regret Has a Golden Hand",
	"Salt, Shame, and Shadows",
	"The Past Is Already Written…",
	"The Flames Lied to Me…",
	"A Song of Ice and Silence…",
	"Patience Is Just Another Word for Revenge",
	"The Viper Never Forgives",
	"Stone Burns Too…",
	"Targaryens Bow to No One",
	"The King Who Knelt and Regretted",
}

var contents = []string{
	"Flames lick at my feet,  \nCrown was mine, I was their queen,  \nNow the dark consumes.  // Rhaenyra’s Last Thought",

	"Dragon’s breath is hot,  \nOne last flight into the storm,  \nSteel meets sky and falls.  // Daemon Targaryen",

	"Golden crown, so bright,  \nYet it crushes bone and flesh,  \nWas this worth the war?  // Aegon II",

	"Whispers in the dark,  \nThe rats are watching me now,  \nDo they know my fate?  // Helaena Targaryen",

	"Should have struck him down,  \nNow the storm is at my door,  \nRegret comes too late.  // Aemond Targaryen",

	"A dream of my sons,  \nA throne to unite them all,  \nWoke to endless war.  // Viserys I",

	"Secrets, webs of lies,  \nA kingdom’s fate in my hands,  \nPower is my god.  // Larys Strong",

	"Steel in trembling hands,  \nHonor bends but does not break,  \nOr so I once thought.  // Criston Cole",

	"Waves crash on my ship,  \nGold and salt mean nothing now,  \nWhat is legacy?  // Corlys Velaryon",

	"Step, step, shadows dance,  \nNot today, death, not today,  \nOne name on my lips.  // Arya Stark",

	"The blade sings its song,  \nJustice is a fleeting wind,  \nCold steel cuts me down.  // Ned Stark",

	"My chains fall away,  \nYet I am still not myself,  \nWho is Theon now?  // Theon Greyjoy",

	"Golden hand feels weak,  \nIs this how my tale must end?  \nDrowning in regret.  // Jaime Lannister",

	"The throne should be mine,  \nSongs will tell of me, I swear,  \nAsh will mark my fall.  // Cersei Lannister",

	"Books and wine still flow,  \nBut wit alone will not save,  \nFools wear golden chains.  // Tyrion Lannister",

	"I was once a king,  \nStorms and swords bowed to my name,  \nNow the wind is still.  // Robert Baratheon",

	"Flames rise in my hands,  \nA god’s will or a fool’s hope?  \nOnly night will tell.  // Melisandre",

	"The past calls my name,  \nBut I walk through time unseen,  \nIs this truly me?  // Bran Stark",

	"Ice and fire clash,  \nThe world bends to silent death,  \nDarkness has no end.  // Night King",

	"Cold winds howl my name,  \nI swore an oath, now I die,  \nFor what cause, I ask?  // Jon Snow",

	"Eyes wide, breath is thin,  \nI see what should not be seen,  \nThe flames lied to me.  // Melisandre",

	"Chains forged from my past,  \nNames I loved, names I betrayed,  \nAll whisper my fate.  // Theon Greyjoy",

	"Sword raised, heart unsure,  \nTo be king or to be free,  \nOne choice, no escape.  // Robb Stark",

	"A knife in the dark,  \nIt sings with a gentle touch,  \nDeath calls me forward.  // Littlefinger",

	"The night is silent,  \nBut I hear them in the wind,  \nThe dead do not sleep.  // Jon Snow",
}

var tags = []string{
	"#HouseOfTheDragon", "#GameOfThrones", "#FireAndBlood", "#IronThrone",
	"#Valyria", "#Dracarys", "#WinterIsComing", "#IceAndFire", "#TheDanceOfDragons",
	"#TargaryenLegacy", "#SwordAndCrown", "#BetrayalAndHonor", "#KingsAndQueens",
	"#TheLastDragon", "#AStormOfSwords", "#DarkWingsDarkWords", "#OathAndDuty",
	"#ShadowsAndFire", "#TheOldGods", "#TheNightIsDark", "#TrialByCombat",
	"#ThePrinceThatWasPromised", "#NoOne", "#TheRedWedding", "#TheKingInTheNorth",
	"#TheMadQueen", "#BloodAndCheese", "#Stormborn", "#WhatIsDeadMayNeverDie",
	"#ThePackSurvives", "#BreakerOfChains", "#TheNightKing", "#TheLannisterDebt",
	"#TheDragonHasThreeHeads", "#UnbowedUnbentUnbroken", "#ChaosIsALadder",
	"#TheGoldenCompany", "#TheRainsOfCastamere", "#AllMenMustDie", "#AFeastForCrows",
}

var comments = []string{
	"Rhaenyra’s Last Thought: ‘It should have been me…’ Yeah, and the BBQ wasn’t part of the plan either. 🔥🐉",

	"Daemon’s Dance: No thoughts, just vibes… until gravity kicked in. 🚀💀",

	"Aegon II’s Crown: Heavy is the head that wears a stolen throne. 👑",

	"Helaena’s Whispers: Ma’am, if the rats are whispering to you, it’s time to log out. 🐀📢",

	"Aemond’s Regret: Maybe don’t start beef with your nephew next time? 👀",

	"Viserys I’s Dream: Should’ve spent less time dreaming, more time parenting. 😬",

	"Larys Strong’s Secrets: We know what your true interest is... Quentin Tarantino style 💅️",

	"Criston Cole’s Honor: Sir Simpington of the Kingsguard strikes again. ⚔️",

	"Corlys Velaryon’s Legacy: Richest man in Westeros, still no peace. Oof. 🏴‍☠️",

	"Arya’s Shadows: Death was all set, and Arya said ‘New phone, who dis?’ ☠️📞",

	"Ned Stark’s Justice: Headstrong? More like just gone. ⚔️🥶",

	"Theon’s Identity Crisis: Prince? Reek? Hero? Bro has had more rebrands than Twitter. 🌀🐺",

	"Jaime’s Redemption: Years of growth, threw it out for his sister. Peak Lannister. 🚮",

	"Cersei’s Fall: ‘They will sing of me forever’—yeah, as a meme. 🏰💥",

	"Tyrion’s Chains: Outdrank, outwitted, outlived… barely. 🍷💀",

	"Robert’s Reflection: Ate himself into irrelevance. Kingship speedrun any% 🍗👑",

	"Melisandre’s Fire: ‘The Lord of Light has a plan’—and it involves *you* taking a massive L. 🔥🤡",

	"Bran’s Time Travel: Biggest plot hole... deux ex machina killed GoT's final season. 🦅",

	"Night King’s End: 8,000 years of prep, undone by a knife toss. Tough. ❄️🗡️",

	"Jon’s Oath: Stabbed, exiled, got zero thanks. Westeros’ worst work contract. 📜🔪",

	"Littlefinger’s Last Words: ‘Chaos is a ladder’\n\nTurns out ladders don’t work when your throat’s cut. 🔪💀",

	"Robb’s War Plan: ‘I got this.’\n\nYeah, so did the Freys. 🩸🔪",
}

func Seed(store store.Storage) {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())

	users := generateUsers(20) // Reduced to 20 for multiple posts per user
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user", user, err)
			return
		}
	}

	posts := generatePosts(50, users) // Each user gets ~2-3 posts
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post", post, err)
			return
		}
	}

	comments := generateComments(100, users, posts) // Each post gets ~2-3 comments
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment", comment, err)
			return
		}
	}

	log.Println("Seeding complete!")
}

func generateUsers(count int) []*store.User {
	users := make([]*store.User, count)
	for i := 0; i < count; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@got.org",
			Password: "password" + strconv.Itoa(i),
		}
	}
	return users
}

func generatePosts(count int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, count)
	for i := 0; i < count; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(count int, users []*store.User, posts []*store.Post) []*store.Comment {
	cmts := make([]*store.Comment, count)

	for i := 0; i < count; i++ {
		post := posts[rand.Intn(len(posts))]
		user := users[rand.Intn(len(users))]

		// Prevent the post author from commenting on their own post
		for user.ID == post.UserID {
			user = users[rand.Intn(len(users))]
		}

		cmts[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cmts
}
