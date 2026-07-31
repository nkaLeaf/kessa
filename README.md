# Kessa



Kessa is a Discord bot built on [discordgo](https://github.com/bwmarrin/discordgo) — a modernized, opinionated take on the OwO Bot–style pet-collecting/progression loop, with its own systems layered on top: turn-based pet battles, prestige progression, an auto-scavenger, a player market, and more. > **Status: early / active development.** Core systems are being built incrementally and numbers, mechanics, and schemas are still subject to change. This README will grow alongside the bot.



---

## Features 

- Progression & Prestige — level up, then prestige through 4 tiers (plus uncapped cosmetic prestige beyond that) for permanent XP bonuses and higher hunter-upgrade caps. Moons, stardust, items, and cosmetics always carry over — prestiging resets levels, not your collection.
- Hunter Upgrades — invest in Stamina, Efficiency, Negotiation, Goggles, and Auto-Filter to shape how scavenging works for you.
- SharkBot (Auto-Scavenger) — send it off to passively collect pets over time, then come back and claim your haul. Timestamp-based, no background jobs.

- Pets — a large roster of collectible pets across 12 rarity tiers, each with its own passive kit (triggers + effects) that gets more complex at higher tiers. Team up to 3 for battle, with randomized stat rolls on first pull.
- Turn-based Battles - physical / mage / true damage types, pet-specific resistances, speed-based turn order, crits and status effects.
- Weapons & Charms — give your pets weapons (with real tradeoffs, not all upside) and wear charms yourself to boost scavenging, chat XP or battle XP.

- Chests & Market — open tiered chests for charms and items, or trade with other players on a global P2P marketplace.

- Contracts — daily, weekly, and community-wide objectives for bonus Moons and XP.

- Gambling - coinflip, blackjack, roulette and a multiplier game like crash.

### Tech stack
 
- **Language:** Go (language)
- **discord lib:** [discordgo](https://github.com/bwmarrin/discordgo)
- **ORM:** [GORM](https://gorm.io/) (SQL database - schema grows feature by feature rather than being fully fixed upfront)# 1. Introduction## Prerequisites 
 
 - [Go](https://go.dev/dl/) (recent stable version) 
 - A SQL database supported by GORM (e.g. PostgreSQL, MySQL, or SQLite) 
 - A [Discord bot application](https://discord.com/developers/applications) and its bot token ### Installation ```bash git clone https://github.com/<your-org>/kessa.git cd kessa go mod download ``` #### Config
 
Copy the example config, and fill in your values:
```bash cp conf.example.json conf.json ``` `conf.json` currently expects at minimum: ```json { "token": "your-discord-bot-token", "prefix": "!" } ```

More config (such as embed colors) will be added as those features land - check `conf.example.json` for latest expected fields.

## Contributing 
 
This project is under active and evolving development. If you want to lend a hand:
 
1. Create an issue to discuss major changes before creating a pull request
2. PRs should be focused - one feature or fix per PR.
3. Follow the current code standards of the repository.
Bug reports and feature requests are welcome on [Issues](../../issues).
 

 
*Numbers, mechanics and schemas described above are subject to change as the bot evolves.*
