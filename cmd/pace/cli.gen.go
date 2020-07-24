// Code generated by oto; DO NOT EDIT.

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
	if len(args) < 2 {
		showHelp(args)
		return nil
	}
	switch args[1] {
	case "help":
		showHelp(args)
		return nil
	case "version":
		fmt.Println(Version)
		return nil
	case "list":
		printList()
		return nil
	case "templates":
		printTemplates()
		return nil

	case "CardsService.CreateCard":
		return CardsServiceCreateCard(ctx, args)

	case "CardsService.DeleteCard":
		return CardsServiceDeleteCard(ctx, args)

	case "CardsService.GetCard":
		return CardsServiceGetCard(ctx, args)

	case "CardsService.PutBackCard":
		return CardsServicePutBackCard(ctx, args)

	case "CardsService.TakeCard":
		return CardsServiceTakeCard(ctx, args)

	case "CardsService.UpdateCard":
		return CardsServiceUpdateCard(ctx, args)

	case "CardsService.UpdateCardStatus":
		return CardsServiceUpdateCardStatus(ctx, args)

	case "CommentsService.AddComment":
		return CommentsServiceAddComment(ctx, args)

	case "CommentsService.DeleteComment":
		return CommentsServiceDeleteComment(ctx, args)

	default:
		fmt.Println("unknown command:", args[1])
		showHelp(args)
	}
	return nil
}

func showHelp(args []string) {
	if len(args) < 3 {
		printUsage()

		fmt.Print("* CardsService")
		commentForCardsService := `CardsService allows you to programmatically manage cards in Pace.`
		if commentForCardsService != "" {
			fmt.Print(" - ", commentForCardsService)
		}
		fmt.Println()

		fmt.Print("* CommentsService")
		commentForCommentsService := `CommentsService allows you to programmatically manage comments in Pace.`
		if commentForCommentsService != "" {
			fmt.Print(" - ", commentForCommentsService)
		}
		fmt.Println()

		fmt.Println()
		fmt.Println("  pace help <service>[.<method>] - print specific help")
		fmt.Println("  pace list - list all services and methods")
		fmt.Println("  pace templates - show copy and paste examples")
		printFlagDefaults(args)
		return
	}
	showHelpFor(args, args[2])
}

func showHelpFor(args []string, service string) {
	switch service {

	case "CardsService":
		fmt.Printf("methods for %s:\n", service)

		fmt.Print("* CardsService.CreateCard")
		commentForCardsServiceCreateCard := `CreateCard creates a new Card.`
		if commentForCardsServiceCreateCard != "" {
			fmt.Print(" - ", commentForCardsServiceCreateCard)
		}
		fmt.Println()

		fmt.Print("* CardsService.DeleteCard")
		commentForCardsServiceDeleteCard := `DeleteCard deletes a card.`
		if commentForCardsServiceDeleteCard != "" {
			fmt.Print(" - ", commentForCardsServiceDeleteCard)
		}
		fmt.Println()

		fmt.Print("* CardsService.GetCard")
		commentForCardsServiceGetCard := `GetCard gets a card.`
		if commentForCardsServiceGetCard != "" {
			fmt.Print(" - ", commentForCardsServiceGetCard)
		}
		fmt.Println()

		fmt.Print("* CardsService.PutBackCard")
		commentForCardsServicePutBackCard := `PutBackCard removes a user from the list of responsbile users.
Undoes TakeCard.`
		if commentForCardsServicePutBackCard != "" {
			fmt.Print(" - ", commentForCardsServicePutBackCard)
		}
		fmt.Println()

		fmt.Print("* CardsService.TakeCard")
		commentForCardsServiceTakeCard := `TakeCard takes responsibility for a card.
Can be undone with PutBackCard.`
		if commentForCardsServiceTakeCard != "" {
			fmt.Print(" - ", commentForCardsServiceTakeCard)
		}
		fmt.Println()

		fmt.Print("* CardsService.UpdateCard")
		commentForCardsServiceUpdateCard := `UpdateCard updates the title and body of the card.`
		if commentForCardsServiceUpdateCard != "" {
			fmt.Print(" - ", commentForCardsServiceUpdateCard)
		}
		fmt.Println()

		fmt.Print("* CardsService.UpdateCardStatus")
		commentForCardsServiceUpdateCardStatus := `UpdateCardStatus updates a card&#39;s status.`
		if commentForCardsServiceUpdateCardStatus != "" {
			fmt.Print(" - ", commentForCardsServiceUpdateCardStatus)
		}
		fmt.Println()

	case "CardsService.CreateCard":
		fmt.Println(`Usage for CardsService.CreateCard

  pace CardsService.CreateCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.CreateCardRequest
		addFlagsForCreateCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.DeleteCard":
		fmt.Println(`Usage for CardsService.DeleteCard

  pace CardsService.DeleteCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.DeleteCardRequest
		addFlagsForDeleteCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.GetCard":
		fmt.Println(`Usage for CardsService.GetCard

  pace CardsService.GetCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.GetCardRequest
		addFlagsForGetCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.PutBackCard":
		fmt.Println(`Usage for CardsService.PutBackCard

  pace CardsService.PutBackCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.PutBackCardRequest
		addFlagsForPutBackCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.TakeCard":
		fmt.Println(`Usage for CardsService.TakeCard

  pace CardsService.TakeCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.TakeCardRequest
		addFlagsForTakeCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.UpdateCard":
		fmt.Println(`Usage for CardsService.UpdateCard

  pace CardsService.UpdateCard [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.UpdateCardRequest
		addFlagsForUpdateCardRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CardsService.UpdateCardStatus":
		fmt.Println(`Usage for CardsService.UpdateCardStatus

  pace CardsService.UpdateCardStatus [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.UpdateCardStatusRequest
		addFlagsForUpdateCardStatusRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CommentsService":
		fmt.Printf("methods for %s:\n", service)

		fmt.Print("* CommentsService.AddComment")
		commentForCommentsServiceAddComment := `AddComment adds a comment.`
		if commentForCommentsServiceAddComment != "" {
			fmt.Print(" - ", commentForCommentsServiceAddComment)
		}
		fmt.Println()

		fmt.Print("* CommentsService.DeleteComment")
		commentForCommentsServiceDeleteComment := `DeleteComment deletes a Comment.`
		if commentForCommentsServiceDeleteComment != "" {
			fmt.Print(" - ", commentForCommentsServiceDeleteComment)
		}
		fmt.Println()

	case "CommentsService.AddComment":
		fmt.Println(`Usage for CommentsService.AddComment

  pace CommentsService.AddComment [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.AddCommentRequest
		addFlagsForAddCommentRequest(flags, "", &request)
		flags.PrintDefaults()

	case "CommentsService.DeleteComment":
		fmt.Println(`Usage for CommentsService.DeleteComment

  pace CommentsService.DeleteComment [flags]

flags:`)
		flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
		// add flags for documentation purposes
		globals := &globalFlags{}
		globals.addFlags(flags)
		var request pace.DeleteCommentRequest
		addFlagsForDeleteCommentRequest(flags, "", &request)
		flags.PrintDefaults()

	default:
		fmt.Println(service, "is not supported")
		showHelp(args)
	}
}

func printUsage() {
	fmt.Printf("%s (%s)\n", Version, ShortSHA)
	fmt.Println(`Usage:
  pace <service>.<method> [args...]`)
	fmt.Println()
}

func CardsServiceCreateCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.CreateCardRequest
	addFlagsForCreateCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
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
	return nil
}

func CardsServiceDeleteCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.DeleteCardRequest
	addFlagsForDeleteCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.DeleteCard(ctx, request)
	if err != nil {
		return err
	}
	printDeleteCardResponse(resp)
	return nil
}

func CardsServiceGetCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.GetCardRequest
	addFlagsForGetCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
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
	return nil
}

func CardsServicePutBackCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.PutBackCardRequest
	addFlagsForPutBackCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.PutBackCard(ctx, request)
	if err != nil {
		return err
	}
	printPutBackCardResponse(resp)
	return nil
}

func CardsServiceTakeCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.TakeCardRequest
	addFlagsForTakeCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.TakeCard(ctx, request)
	if err != nil {
		return err
	}
	printTakeCardResponse(resp)
	return nil
}

func CardsServiceUpdateCard(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.UpdateCardRequest
	addFlagsForUpdateCardRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.UpdateCard(ctx, request)
	if err != nil {
		return err
	}
	printUpdateCardResponse(resp)
	return nil
}

func CardsServiceUpdateCardStatus(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.UpdateCardStatusRequest
	addFlagsForUpdateCardStatusRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCardsService(client)
	resp, err := service.UpdateCardStatus(ctx, request)
	if err != nil {
		return err
	}
	printUpdateCardStatusResponse(resp)
	return nil
}

func CommentsServiceAddComment(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.AddCommentRequest
	addFlagsForAddCommentRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
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
	return nil
}

func CommentsServiceDeleteComment(ctx context.Context, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	var request pace.DeleteCommentRequest
	addFlagsForDeleteCommentRequest(flags, "", &request)
	if err := flags.Parse(args[2:]); err != nil {
		return err
	}
	if globals.apikey == "" {
		globals.apikey = os.Getenv("PACE_API_KEY")
	}
	if globals.apikey == "" {
		return errors.New("missing api key (use -apikey flag or PACE_API_KEY env var)")
	}
	if globals.secret == "" {
		globals.secret = os.Getenv("PACE_API_SECRET")
	}
	if globals.secret == "" {
		return errors.New("missing api secret (use -secret flag or PACE_API_SECRET env var)")
	}
	client := pace.New(globals.apikey, globals.secret)
	client.RemoteHost = globals.host
	if globals.debug {
		client.Debug = func(s string) {
			fmt.Println(s)
		}
	}
	service := pace.NewCommentsService(client)
	resp, err := service.DeleteComment(ctx, request)
	if err != nil {
		return err
	}
	printDeleteCommentResponse(resp)
	return nil
}

func printFlagDefaults(args []string) {
	if len(args) == 0 {
		// avoid empty arg panics
		args = []string{"pace"}
	}
	fmt.Println()
	fmt.Println("Flags:")
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	globals := &globalFlags{}
	globals.addFlags(flags)
	flags.PrintDefaults()
}

func printList() {

	fmt.Printf("CardsService.CreateCard")
	commentCardsServiceCreateCard := `CreateCard creates a new Card.`
	if len(commentCardsServiceCreateCard) > 0 {
		fmt.Printf(" - %s", commentCardsServiceCreateCard)
	}
	fmt.Println()

	fmt.Printf("CardsService.DeleteCard")
	commentCardsServiceDeleteCard := `DeleteCard deletes a card.`
	if len(commentCardsServiceDeleteCard) > 0 {
		fmt.Printf(" - %s", commentCardsServiceDeleteCard)
	}
	fmt.Println()

	fmt.Printf("CardsService.GetCard")
	commentCardsServiceGetCard := `GetCard gets a card.`
	if len(commentCardsServiceGetCard) > 0 {
		fmt.Printf(" - %s", commentCardsServiceGetCard)
	}
	fmt.Println()

	fmt.Printf("CardsService.PutBackCard")
	commentCardsServicePutBackCard := `PutBackCard removes a user from the list of responsbile users.
Undoes TakeCard.`
	if len(commentCardsServicePutBackCard) > 0 {
		fmt.Printf(" - %s", commentCardsServicePutBackCard)
	}
	fmt.Println()

	fmt.Printf("CardsService.TakeCard")
	commentCardsServiceTakeCard := `TakeCard takes responsibility for a card.
Can be undone with PutBackCard.`
	if len(commentCardsServiceTakeCard) > 0 {
		fmt.Printf(" - %s", commentCardsServiceTakeCard)
	}
	fmt.Println()

	fmt.Printf("CardsService.UpdateCard")
	commentCardsServiceUpdateCard := `UpdateCard updates the title and body of the card.`
	if len(commentCardsServiceUpdateCard) > 0 {
		fmt.Printf(" - %s", commentCardsServiceUpdateCard)
	}
	fmt.Println()

	fmt.Printf("CardsService.UpdateCardStatus")
	commentCardsServiceUpdateCardStatus := `UpdateCardStatus updates a card&#39;s status.`
	if len(commentCardsServiceUpdateCardStatus) > 0 {
		fmt.Printf(" - %s", commentCardsServiceUpdateCardStatus)
	}
	fmt.Println()

	fmt.Printf("CommentsService.AddComment")
	commentCommentsServiceAddComment := `AddComment adds a comment.`
	if len(commentCommentsServiceAddComment) > 0 {
		fmt.Printf(" - %s", commentCommentsServiceAddComment)
	}
	fmt.Println()

	fmt.Printf("CommentsService.DeleteComment")
	commentCommentsServiceDeleteComment := `DeleteComment deletes a Comment.`
	if len(commentCommentsServiceDeleteComment) > 0 {
		fmt.Printf(" - %s", commentCommentsServiceDeleteComment)
	}
	fmt.Println()

}

func printTemplates() {

	fmt.Printf("pace CardsService.CreateCard ")
	printArgslistCreateCardRequest()
	fmt.Println()

	fmt.Printf("pace CardsService.DeleteCard ")
	printArgslistDeleteCardRequest()
	fmt.Println()

	fmt.Printf("pace CardsService.GetCard ")
	printArgslistGetCardRequest()
	fmt.Println()

	fmt.Printf("pace CardsService.PutBackCard ")
	printArgslistPutBackCardRequest()
	fmt.Println()

	fmt.Printf("pace CardsService.TakeCard ")
	printArgslistTakeCardRequest()
	fmt.Println()

	fmt.Printf("pace CardsService.UpdateCard ")
	printArgslistUpdateCardRequest()
	fmt.Println()

	fmt.Printf("pace CardsService.UpdateCardStatus ")
	printArgslistUpdateCardStatusRequest()
	fmt.Println()

	fmt.Printf("pace CommentsService.AddComment ")
	printArgslistAddCommentRequest()
	fmt.Println()

	fmt.Printf("pace CommentsService.DeleteComment ")
	printArgslistDeleteCommentRequest()
	fmt.Println()

}

func addFlagsForAddCommentRequest(flags *flag.FlagSet, prefix string, v *pace.AddCommentRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of the org.`)

	flags.StringVar(&v.TargetKind, prefix+"targetKind", "", `TargetKind is the kind of item this comment is for.
Can be &#34;card&#34;, &#34;message&#34;, or &#34;showcase&#34;.`)

	flags.StringVar(&v.TargetID, prefix+"targetID", "", `TargetID is the ID of the target.`)

	flags.StringVar(&v.Body, prefix+"body", "", `Body is the markdown body of the comment.`)

}

func printAddCommentRequest(v *pace.AddCommentRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistAddCommentRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-targetKind= ")

	fmt.Print("-targetID= ")

	fmt.Print("-body= ")

}

func addFlagsForPerson(flags *flag.FlagSet, prefix string, v *pace.Person) {

	flags.StringVar(&v.ID, prefix+"id", "", `ID is the ID of the Person.`)

	flags.StringVar(&v.Username, prefix+"username", "", `Username is the Person&#39;s username within the org.`)

	flags.StringVar(&v.Name, prefix+"name", "", `Name is the name of the Person.`)

	flags.StringVar(&v.PhotoURL, prefix+"photoURL", "", `PhotoURL is the URL of a picture of this Person.`)

}

func printPerson(v *pace.Person) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistPerson() {

	fmt.Print("-id= ")

	fmt.Print("-username= ")

	fmt.Print("-name= ")

	fmt.Print("-photoURL= ")

}

func addFlagsForComment(flags *flag.FlagSet, prefix string, v *pace.Comment) {

	flags.StringVar(&v.ID, prefix+"id", "", `ID is the ID of the comment.`)

	flags.StringVar(&v.CTime, prefix+"cTime", "", `CTime is the time this was created.`)

	flags.StringVar(&v.MTime, prefix+"mTime", "", `MTime is the time this comment was last modified.`)

	flags.StringVar(&v.Body, prefix+"body", "", `Body is the markdown body of the comment.`)

	flags.StringVar(&v.BodyHTML, prefix+"bodyHTML", "", `BodyHTML is the HTML formatted body of the comment.`)

	addFlagsForPerson(flags, "author.", &v.Author)

}

func printComment(v *pace.Comment) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistComment() {

	fmt.Print("-id= ")

	fmt.Print("-cTime= ")

	fmt.Print("-mTime= ")

	fmt.Print("-body= ")

	fmt.Print("-bodyHTML= ")

	fmt.Print("-author= ")

}

func addFlagsForAddCommentResponse(flags *flag.FlagSet, prefix string, v *pace.AddCommentResponse) {

	addFlagsForComment(flags, "comment.", &v.Comment)

	// skipping Error field (handled by Go client)

}

func printAddCommentResponse(v *pace.AddCommentResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistAddCommentResponse() {

	fmt.Print("-comment= ")

	fmt.Print("-error= ")

}

func addFlagsForFile(flags *flag.FlagSet, prefix string, v *pace.File) {

	flags.StringVar(&v.ID, prefix+"id", "", `ID is the identifier for this file.`)

	flags.StringVar(&v.CTime, prefix+"cTime", "", `CTime is the time the file was uploaded.`)

	flags.StringVar(&v.Name, prefix+"name", "", `Name is the name of the file.`)

	flags.StringVar(&v.Path, prefix+"path", "", `Path is the path of the file.`)

	flags.StringVar(&v.ContentType, prefix+"contentType", "", `ContentType is the type of the file.`)

	flags.StringVar(&v.FileType, prefix+"fileType", "", `FileType is the type of file.
Can be &#34;file&#34;, &#34;video&#34;, &#34;image&#34;, &#34;audio&#34; or &#34;screenshare&#34;.`)

	flags.StringVar(&v.DownloadURL, prefix+"downloadURL", "", `DownloadURL URL which can be used to get the file.`)

	flags.StringVar(&v.ThumbnailURL, prefix+"thumbnailURL", "", `ThumbnailURL is an optional thumbnail URL for this file.`)

	addFlagsForPerson(flags, "author.", &v.Author)

}

func printFile(v *pace.File) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistFile() {

	fmt.Print("-id= ")

	fmt.Print("-cTime= ")

	fmt.Print("-name= ")

	fmt.Print("-path= ")

	fmt.Print("-contentType= ")

	fmt.Print("-fileType= ")

	fmt.Print("-size= ")

	fmt.Print("-downloadURL= ")

	fmt.Print("-thumbnailURL= ")

	fmt.Print("-author= ")

}

func addFlagsForCard(flags *flag.FlagSet, prefix string, v *pace.Card) {

	flags.StringVar(&v.ID, prefix+"id", "", `ID is the unique ID of the card within the org.`)

	flags.StringVar(&v.CTime, prefix+"cTime", "", `CTime is the time this was created.`)

	flags.StringVar(&v.MTime, prefix+"mTime", "", `MTime is the time this comment was last modified.`)

	flags.StringVar(&v.TeamID, prefix+"teamID", "", `TeamID is the ID of the team that this card belongs to.`)

	flags.StringVar(&v.Slug, prefix+"slug", "", `Slug is the URL slug for this card.`)

	flags.StringVar(&v.Title, prefix+"title", "", `Title is the title of the card.`)

	flags.StringVar(&v.Status, prefix+"status", "", `Status is the current status of the card.`)

	addFlagsForPerson(flags, "author.", &v.Author)

	flags.StringVar(&v.Body, prefix+"body", "", `Body is the markdown body of this card.`)

	flags.StringVar(&v.BodyHTML, prefix+"bodyHTML", "", `BodyHTML is the HTML rendering of the body of this card.`)

	// []string not supported yet

	// []Person not supported yet

	// []File not supported yet

}

func printCard(v *pace.Card) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistCard() {

	fmt.Print("-id= ")

	fmt.Print("-cTime= ")

	fmt.Print("-mTime= ")

	fmt.Print("-teamID= ")

	fmt.Print("-slug= ")

	fmt.Print("-title= ")

	fmt.Print("-status= ")

	fmt.Print("-author= ")

	fmt.Print("-body= ")

	fmt.Print("-bodyHTML= ")

	fmt.Print("-tags= ")

	fmt.Print("-takenByCurrentUser= ")

	fmt.Print("-takenByPeople= ")

	fmt.Print("-files= ")

}

func addFlagsForCreateCardRequest(flags *flag.FlagSet, prefix string, v *pace.CreateCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the org ID in which to create the card.`)

	flags.StringVar(&v.TeamID, prefix+"teamID", "", `TeamID is the team ID in which to create the card.`)

	flags.StringVar(&v.Title, prefix+"title", "", `Title is the title of the card.`)

	flags.StringVar(&v.ParentTargetKind, prefix+"parentTargetKind", "", `ParentTargetKind is the kind of target to relate this card to (e.g. &#34;card&#34;, &#34;message&#34;, or &#34;showcase&#34;)`)

	flags.StringVar(&v.ParentTargetID, prefix+"parentTargetID", "", `ParentTargetID is the ID of the item to relate this new card to.`)

}

func printCreateCardRequest(v *pace.CreateCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistCreateCardRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-teamID= ")

	fmt.Print("-title= ")

	fmt.Print("-parentTargetKind= ")

	fmt.Print("-parentTargetID= ")

}

func addFlagsForCreateCardResponse(flags *flag.FlagSet, prefix string, v *pace.CreateCardResponse) {

	addFlagsForCard(flags, "card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printCreateCardResponse(v *pace.CreateCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistCreateCardResponse() {

	fmt.Print("-card= ")

	fmt.Print("-error= ")

}

func addFlagsForDeleteCardRequest(flags *flag.FlagSet, prefix string, v *pace.DeleteCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of your org.`)

	flags.StringVar(&v.CardID, prefix+"cardID", "", `CardID is the ID of the card to delete.`)

}

func printDeleteCardRequest(v *pace.DeleteCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistDeleteCardRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-cardID= ")

}

func addFlagsForDeleteCardResponse(flags *flag.FlagSet, prefix string, v *pace.DeleteCardResponse) {

	// skipping Error field (handled by Go client)

}

func printDeleteCardResponse(v *pace.DeleteCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistDeleteCardResponse() {

	fmt.Print("-error= ")

}

func addFlagsForGetCardRequest(flags *flag.FlagSet, prefix string, v *pace.GetCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of the org.`)

	flags.StringVar(&v.CardID, prefix+"cardID", "", `CardID is the ID of the card to get.`)

}

func printGetCardRequest(v *pace.GetCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistGetCardRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-cardID= ")

}

func addFlagsForGetCardResponse(flags *flag.FlagSet, prefix string, v *pace.GetCardResponse) {

	addFlagsForCard(flags, "card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printGetCardResponse(v *pace.GetCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistGetCardResponse() {

	fmt.Print("-card= ")

	fmt.Print("-error= ")

}

func addFlagsForPutBackCardRequest(flags *flag.FlagSet, prefix string, v *pace.PutBackCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of your org.`)

	flags.StringVar(&v.CardID, prefix+"cardID", "", `CardID is the ID of the card to unassign yourself from.`)

}

func printPutBackCardRequest(v *pace.PutBackCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistPutBackCardRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-cardID= ")

}

func addFlagsForPutBackCardResponse(flags *flag.FlagSet, prefix string, v *pace.PutBackCardResponse) {

	addFlagsForCard(flags, "card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printPutBackCardResponse(v *pace.PutBackCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistPutBackCardResponse() {

	fmt.Print("-card= ")

	fmt.Print("-error= ")

}

func addFlagsForTakeCardRequest(flags *flag.FlagSet, prefix string, v *pace.TakeCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of your org.`)

	flags.StringVar(&v.CardID, prefix+"cardID", "", `CardID is the ID of the card to assign yourself to.`)

}

func printTakeCardRequest(v *pace.TakeCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistTakeCardRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-cardID= ")

}

func addFlagsForTakeCardResponse(flags *flag.FlagSet, prefix string, v *pace.TakeCardResponse) {

	addFlagsForCard(flags, "card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printTakeCardResponse(v *pace.TakeCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistTakeCardResponse() {

	fmt.Print("-card= ")

	fmt.Print("-error= ")

}

func addFlagsForUpdateCardRequest(flags *flag.FlagSet, prefix string, v *pace.UpdateCardRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of the org.`)

	flags.StringVar(&v.CardID, prefix+"cardID", "", `CardID is the ID of the card to update.`)

	flags.StringVar(&v.Title, prefix+"title", "", `Title is the new title for the card.`)

	flags.StringVar(&v.Body, prefix+"body", "", `Body is the new markdown body for the card.`)

}

func printUpdateCardRequest(v *pace.UpdateCardRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistUpdateCardRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-cardID= ")

	fmt.Print("-title= ")

	fmt.Print("-body= ")

}

func addFlagsForUpdateCardResponse(flags *flag.FlagSet, prefix string, v *pace.UpdateCardResponse) {

	addFlagsForCard(flags, "card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printUpdateCardResponse(v *pace.UpdateCardResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistUpdateCardResponse() {

	fmt.Print("-card= ")

	fmt.Print("-error= ")

}

func addFlagsForUpdateCardStatusRequest(flags *flag.FlagSet, prefix string, v *pace.UpdateCardStatusRequest) {

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of the org.`)

	flags.StringVar(&v.CardID, prefix+"cardID", "", `CardID is the ID number of the card.`)

	flags.StringVar(&v.Status, prefix+"status", "", `Status is the new status of the card.
Valid strings are &#34;future&#34;, &#34;next&#34;, &#34;progress&#34;, &#34;done&#34;.`)

}

func printUpdateCardStatusRequest(v *pace.UpdateCardStatusRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistUpdateCardStatusRequest() {

	fmt.Print("-orgID= ")

	fmt.Print("-cardID= ")

	fmt.Print("-status= ")

}

func addFlagsForUpdateCardStatusResponse(flags *flag.FlagSet, prefix string, v *pace.UpdateCardStatusResponse) {

	addFlagsForCard(flags, "card.", &v.Card)

	// skipping Error field (handled by Go client)

}

func printUpdateCardStatusResponse(v *pace.UpdateCardStatusResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistUpdateCardStatusResponse() {

	fmt.Print("-card= ")

	fmt.Print("-error= ")

}

func addFlagsForDeleteCommentRequest(flags *flag.FlagSet, prefix string, v *pace.DeleteCommentRequest) {

	flags.StringVar(&v.ID, prefix+"id", "", `ID is the ID of the comment to delete.`)

	flags.StringVar(&v.OrgID, prefix+"orgID", "", `OrgID is the ID of your org.`)

	flags.StringVar(&v.TargetKind, prefix+"targetKind", "", `TargetKind is the kind of target this Comment was made on.
Can be &#34;card&#34;, &#34;message&#34;, or &#34;showcase&#34;.
Used to help identify the Comment.`)

	flags.StringVar(&v.TargetID, prefix+"targetID", "", `TargetID is the ID of the target.
Used to help identify the Comment.`)

}

func printDeleteCommentRequest(v *pace.DeleteCommentRequest) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistDeleteCommentRequest() {

	fmt.Print("-id= ")

	fmt.Print("-orgID= ")

	fmt.Print("-targetKind= ")

	fmt.Print("-targetID= ")

}

func addFlagsForDeleteCommentResponse(flags *flag.FlagSet, prefix string, v *pace.DeleteCommentResponse) {

	// skipping Error field (handled by Go client)

}

func printDeleteCommentResponse(v *pace.DeleteCommentResponse) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("failed to marshal response:", err)
		return
	}
	fmt.Println(string(b))
}

func printArgslistDeleteCommentResponse() {

	fmt.Print("-error= ")

}

type globalFlags struct {
	apikey string
	secret string
	host   string
	debug  bool
	silent bool
}

func (c *globalFlags) addFlags(flags *flag.FlagSet) {
	flags.StringVar(&c.apikey, "apikey", "", "Pace API Key (env var: PACE_API_KEY)")
	flags.StringVar(&c.secret, "secret", "", "Pace API secret (env var: PACE_API_SECRET)")
	flags.StringVar(&c.host, "host", "https://pace.dev", "Pace remote host")
	flags.BoolVar(&c.debug, "debug", false, "prints debug information")
	flags.BoolVar(&c.silent, "silent", false, "prints no output for successful requests")
}
