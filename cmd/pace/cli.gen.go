package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/pacedotdev/pace"
)

func runCommand(ctx context.Context, args []string) error {
	if len(args) < 2 || args[1] == "help" {
		showHelp(args)
		return nil
	}
	switch args[1] {

	case "CardsService.CreateCard":
		return CardsServiceCreateCard(ctx, args)

	case "CardsService.GetCard":
		return CardsServiceGetCard(ctx, args)

	case "CommentsService.AddComment":
		return CommentsServiceAddComment(ctx, args)

	default:
		fmt.Println("unknown command:", args[1])
		showHelp(nil)
	}
	return nil
}

func showHelp(args []string) {
	if len(args) < 3 {
		printUsage()
		fmt.Println("services:")

		fmt.Println("- CardsService")

		fmt.Println("- CommentsService")

		fmt.Println("for more info: pace help <service>")
		return
	}
	showHelpFor(args, args[2])
}

func showHelpFor(args []string, service string) {
	switch service {

	case "CardsService":
		fmt.Printf("methods for %s:\n", service)

		fmt.Println("- CardsService.CreateCard")

		fmt.Println("- CardsService.GetCard")

	case "CardsService.CreateCard":
		fmt.Println(`usage for CardsService.CreateCard

  pace CardsService.CreateCard [flags]`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		var request pace.CreateCardRequest
		addFlagsForCreateCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.GetCard":
		fmt.Println(`usage for CardsService.GetCard

  pace CardsService.GetCard [flags]`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		var request pace.GetCardRequest
		addFlagsForGetCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CommentsService":
		fmt.Printf("methods for %s:\n", service)

		fmt.Println("- CommentsService.AddComment")

	case "CommentsService.AddComment":
		fmt.Println(`usage for CommentsService.AddComment

  pace CommentsService.AddComment [flags]`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		var request pace.AddCommentRequest
		addFlagsForAddCommentRequest(flags, "", &request)
		flags.PrintDefaults()

	default:
		fmt.Println(service, "is not supported")
		showHelp(nil)
	}
}

func printUsage() {
	fmt.Println(`usage: pace <service>.<method> [args...]`)
}

func CardsServiceCreateCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		apikey string
		host   string
		debug  bool
	)
	flags.StringVar(&apikey, "apikey", "", "Pace API Key")
	flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&debug, "debug", false, "prints debug information")
	var request pace.CreateCardRequest
	addFlagsForCreateCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if apikey == "" {
		apikey = os.Getenv("PACE_API_KEY")
	}
	if apikey == "" {
		return errors.New("missing apikey")
	}
	client := pace.New(apikey)
	client.RemoteHost = host
	if debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.CreateCard(ctx, request)
	if err != nil {
		return err
	}
	printCreateCardResponse(resp)
	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

func CardsServiceGetCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		apikey string
		host   string
		debug  bool
	)
	flags.StringVar(&apikey, "apikey", "", "Pace API Key")
	flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&debug, "debug", false, "prints debug information")
	var request pace.GetCardRequest
	addFlagsForGetCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if apikey == "" {
		apikey = os.Getenv("PACE_API_KEY")
	}
	if apikey == "" {
		return errors.New("missing apikey")
	}
	client := pace.New(apikey)
	client.RemoteHost = host
	if debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.GetCard(ctx, request)
	if err != nil {
		return err
	}
	printGetCardResponse(resp)
	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

func CommentsServiceAddComment(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	var (
		apikey string
		host   string
		debug  bool
	)
	flags.StringVar(&apikey, "apikey", "", "Pace API Key")
	flags.StringVar(&host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&debug, "debug", false, "prints debug information")
	var request pace.AddCommentRequest
	addFlagsForAddCommentRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if apikey == "" {
		apikey = os.Getenv("PACE_API_KEY")
	}
	if apikey == "" {
		return errors.New("missing apikey")
	}
	client := pace.New(apikey)
	client.RemoteHost = host
	if debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCommentsService(client)
	resp, err := service.AddComment(ctx, request)
	if err != nil {
		return err
	}
	printAddCommentResponse(resp)
	if resp.Error != "" {
		return errors.New(resp.Error)
	}
	return nil
}

func addFlagsForAddCommentRequest(flags *flag.FlagSet, prefix string, v *pace.AddCommentRequest) {

	flags.StringVar(&v.OrgID, prefix+"OrgID", "", "")

	flags.StringVar(&v.TargetKind, prefix+"TargetKind", "", "")

	flags.StringVar(&v.TargetID, prefix+"TargetID", "", "")

	flags.StringVar(&v.Body, prefix+"Body", "", "")

}

func printAddCommentRequest(v *pace.AddCommentRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForPerson(flags *flag.FlagSet, prefix string, v *pace.Person) {

	flags.StringVar(&v.ID, prefix+"ID", "", "")

	flags.StringVar(&v.Username, prefix+"Username", "", "")

	flags.StringVar(&v.Name, prefix+"Name", "", "")

	flags.StringVar(&v.PhotoURL, prefix+"PhotoURL", "", "")

}

func printPerson(v *pace.Person) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForFile(flags *flag.FlagSet, prefix string, v *pace.File) {

	flags.StringVar(&v.ID, prefix+"ID", "", "")

	flags.StringVar(&v.CTime, prefix+"CTime", "", "")

	flags.StringVar(&v.Name, prefix+"Name", "", "")

	flags.StringVar(&v.Path, prefix+"Path", "", "")

	flags.StringVar(&v.ContentType, prefix+"ContentType", "", "")

	flags.StringVar(&v.FileType, prefix+"FileType", "", "")

	flags.StringVar(&v.DownloadURL, prefix+"DownloadURL", "", "")

	flags.StringVar(&v.ThumbnailURL, prefix+"ThumbnailURL", "", "")

	addFlagsForPerson(flags, "Author.", &v.Author)

}

func printFile(v *pace.File) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForComment(flags *flag.FlagSet, prefix string, v *pace.Comment) {

	flags.StringVar(&v.ID, prefix+"ID", "", "")

	flags.StringVar(&v.CTime, prefix+"CTime", "", "")

	flags.StringVar(&v.MTime, prefix+"MTime", "", "")

	flags.StringVar(&v.Body, prefix+"Body", "", "")

	flags.StringVar(&v.BodyHTML, prefix+"BodyHTML", "", "")

	addFlagsForPerson(flags, "Author.", &v.Author)

	// []File not supported yet

}

func printComment(v *pace.Comment) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForAddCommentResponse(flags *flag.FlagSet, prefix string, v *pace.AddCommentResponse) {

	addFlagsForComment(flags, "Comment.", &v.Comment)

	flags.StringVar(&v.Error, prefix+"Error", "", "")

}

func printAddCommentResponse(v *pace.AddCommentResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForRelatedCardsSummary(flags *flag.FlagSet, prefix string, v *pace.RelatedCardsSummary) {

}

func printRelatedCardsSummary(v *pace.RelatedCardsSummary) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForCard(flags *flag.FlagSet, prefix string, v *pace.Card) {

	flags.StringVar(&v.ID, prefix+"ID", "", "")

	flags.StringVar(&v.CTime, prefix+"CTime", "", "")

	flags.StringVar(&v.MTime, prefix+"MTime", "", "")

	flags.StringVar(&v.TeamID, prefix+"TeamID", "", "")

	flags.StringVar(&v.Slug, prefix+"Slug", "", "")

	flags.StringVar(&v.Title, prefix+"Title", "", "")

	flags.StringVar(&v.Status, prefix+"Status", "", "")

	addFlagsForPerson(flags, "Author.", &v.Author)

	flags.StringVar(&v.Body, prefix+"Body", "", "")

	flags.StringVar(&v.BodyHTML, prefix+"BodyHTML", "", "")

	// []string not supported yet

	// []Person not supported yet

	// []File not supported yet

	addFlagsForRelatedCardsSummary(flags, "RelatedCardsSummary.", &v.RelatedCardsSummary)

}

func printCard(v *pace.Card) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForCreateCardRequest(flags *flag.FlagSet, prefix string, v *pace.CreateCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"OrgID", "", "")

	flags.StringVar(&v.TeamID, prefix+"TeamID", "", "")

	flags.StringVar(&v.Title, prefix+"Title", "", "")

	flags.StringVar(&v.ParentTargetKind, prefix+"ParentTargetKind", "", "")

	flags.StringVar(&v.ParentTargetID, prefix+"ParentTargetID", "", "")

}

func printCreateCardRequest(v *pace.CreateCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForCreateCardResponse(flags *flag.FlagSet, prefix string, v *pace.CreateCardResponse) {

	addFlagsForCard(flags, "Card.", &v.Card)

	flags.StringVar(&v.Error, prefix+"Error", "", "")

}

func printCreateCardResponse(v *pace.CreateCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForGetCardRequest(flags *flag.FlagSet, prefix string, v *pace.GetCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"OrgID", "", "")

	flags.StringVar(&v.CardID, prefix+"CardID", "", "")

}

func printGetCardRequest(v *pace.GetCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func addFlagsForGetCardResponse(flags *flag.FlagSet, prefix string, v *pace.GetCardResponse) {

	addFlagsForCard(flags, "Card.", &v.Card)

	flags.StringVar(&v.Error, prefix+"Error", "", "")

}

func printGetCardResponse(v *pace.GetCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}
