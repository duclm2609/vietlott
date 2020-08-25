package domain

import (
	"fmt"
	"strings"
)

const BlockTypeSection = "section"
const TextTypeMarkdown = "mrkdwn"
const DateLayout = "02/01/2006"

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Block struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type SlackMessage struct {
	Blocks []Block `json:"blocks"`
}

func mapDrawIdText(drawId string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":fire: Kỳ quay: *%s*", drawId),
	}
}

func mapDrawDateText(drawDate string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":date: Ngày quay thưởng: *%s*", drawDate),
	}
}

func mapJackpotPrize(prize string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":dollar: Giá trị Jackpot: *%s* VNĐ", prize),
	}
}

func mapJackpotNumber(prize Jackpot) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":slot_machine: Jackpot: *%s*", prize),
	}
}

func mapNumberOfJackpot(number string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":medal: Số lượng giải Jackpot: *%s*", number),
	}
}

func MapFrom(mega Mega645Result) SlackMessage {
	return SlackMessage{
		Blocks: []Block{
			{
				Type: BlockTypeSection,
				Text: mapDrawIdText(mega.DrawId),
			},
			{
				Type: BlockTypeSection,
				Text: mapDrawDateText(mega.DrawDate.Format(DateLayout)),
			},
			{
				Type: BlockTypeSection,
				Text: mapJackpotPrize(mega.Prize),
			},
			{
				Type: BlockTypeSection,
				Text: mapJackpotNumber(mega.Jackpot),
			},
			{
				Type: BlockTypeSection,
				Text: mapNumberOfJackpot(mega.Winner),
			},
		},
	}
}

func MapFromCompareResult(result Mega645CompareResult) SlackMessage {
	var blocks []Block
	if result.JackpotPrize[0] != 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf(":tada: *Trúng Jackpot rồiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii!!!!!!!!!!!!!!!!!!!*\n%+v", result.JackpotPrize),
			},
		})
	}
	if len(result.FirstPrize) > 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf("Giải nhất: %d\n%s", len(result.FirstPrize), printTicketList(result.FirstPrize)),
			},
		})
	}
	if len(result.SecondPrize) > 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf("Giải nhì: %d\n%s", len(result.SecondPrize), printTicketList(result.SecondPrize)),
			},
		})
	}
	if len(result.ThirdPrize) > 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf("Giải ba: %d\n%s", len(result.ThirdPrize), printTicketList(result.ThirdPrize)),
			},
		})
	}
	if len(blocks) == 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: "Chưa trúng phát nào, chúc may mắn lần sau nhé ông bạn!",
			},
		})
	}
	return SlackMessage{Blocks: blocks}
}

func MapFromTicketList(tickets [][]int) SlackMessage {
	var blocks []Block
	blocks = append(blocks, Block{
		Type: BlockTypeSection,
		Text: Text{
			Type: TextTypeMarkdown,
			Text: ":heart: dãy số may mắn ngày hôm nay :heart:",
		},
	})
	var b strings.Builder
	for _, item := range tickets {
		_, _ = fmt.Fprintf(&b, "%d\n", item)

	}
	blocks = append(blocks, Block{
		Type: BlockTypeSection,
		Text: Text{
			Type: TextTypeMarkdown,
			Text: fmt.Sprintf("%s", b.String()),
		},
	})

	return SlackMessage{Blocks: blocks}
}

func printTicketList(tickets []Ticket) string {
	var b strings.Builder
	for _, t := range tickets {
		_, _ = fmt.Fprintf(&b, "%v\n", t)
	}
	return b.String()
}
