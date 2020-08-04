package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/dgwidgets"
	"github.com/cfi2017/sheep-bot/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

func colorCommandFactory(session *discordgo.Session, event *discordgo.MessageCreate) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "color",
		Aliases: []string{"c", "cl", "colour", "ch"},
		Short:   "Change your role color.",
		Run: func(cmd *cobra.Command, args []string) {
			var color string
			if len(args) == 0 {
				color = util.RandomColor()
			} else {
				// parse color from args or get saved color
			}

			if !util.IsColorValid(color) {
				cmd.PrintErrln("That color is not valid :(")
				return
			}

			color = util.MustParseColor(color)
			if color == "#000000" {
				color = "#000011"
			}

			numericColor, _ := strconv.ParseInt(color[1:], 16, 32)
			embed := &discordgo.MessageEmbed{
				Title:  fmt.Sprintf("Color %s", strings.ToUpper(color)),
				Color:  int(numericColor),
				Footer: &discordgo.MessageEmbedFooter{Text: util.MustColorToRGB(color)},
				Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://%s/%s", viper.GetString("colors.image_endpoint"), color[1:])},
			}

			w := dgwidgets.NewWidget(session, event.ChannelID, embed)
			_ = w.Handle(":white_check_mark:", func(widget *dgwidgets.Widget, reaction *discordgo.MessageReaction) {

			})
			_ = w.Handle(":x:", func(widget *dgwidgets.Widget, reaction *discordgo.MessageReaction) {
				_ = session.ChannelMessageDelete(event.ChannelID, w.Message.ID)
				_ = w.Close()
			})
			_ = w.Handle(":twisted_rightwards_arrows:", func(widget *dgwidgets.Widget, reaction *discordgo.MessageReaction) {

			})
			w.DeleteOnTimeout = true
			w.Timeout = 90 * time.Second
			w.UserWhitelist = []string{event.Author.ID}
			err := w.Spawn()
			if err != nil {
				cmd.PrintErrln("Unknown error trying to respond to command.")
			}

		},
	}

	return cmd
}
