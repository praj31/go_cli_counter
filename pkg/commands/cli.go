package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/praj31/cli_counter/pkg/cache"
	"github.com/thatisuday/commando"
)

func Launch() {
	rdb := cache.GetRedisClient()

	commando.
		SetExecutableName("counter").
		SetVersion("1.0.0").
		SetDescription("Utility counter storage tool backed by Redis.")

	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Println("-- COUNTER --")
			fmt.Println("Utility counter storage tool backed by Redis.")
			fmt.Println("Run `counter --help` for usage details.")
		})
	commando.
		Register("list").
		SetDescription("The list command is used to display all the counters in the list.").
		SetShortDescription("displays counter list").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			rdb.ListCounters()
		})

	commando.
		Register("add").
		SetDescription("The add command is used to add a new counter to the list.").
		SetShortDescription("adds a new counter").
		AddArgument("name", "Name of the counter", "DefaultCounter").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			rdb.AddCounter(args["name"].Value)
		})

	commando.
		Register("remove").
		SetDescription("The remove command is used to remove a counter from the list.").
		SetShortDescription("removes a counter").
		AddArgument("name", "Name of the counter", "DefaultCounter").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			rdb.RemoveCounter(args["name"].Value)
		})

	commando.
		Register("incr").
		SetDescription("The incr command is used to increment a counter by 1.").
		SetShortDescription("increments a counter by 1").
		AddArgument("name", "Name of the counter", "DefaultCounter").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			rdb.IncrementCounter(args["name"].Value)
		})

	commando.
		Register("decr").
		SetDescription("The decr command is used to decrement a counter by 1.").
		SetShortDescription("decrements a counter by 1").
		AddArgument("name", "Name of the counter", "DefaultCounter").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			rdb.DecrementCounter(args["name"].Value)
		})

	commando.
		Register("get").
		SetDescription("The get command is used to obtain a counter's value from the list.").
		SetShortDescription("gets a counter's value").
		AddArgument("name", "Name of the counter", "DefaultCounter").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			rdb.GetCounter(args["name"].Value)
		})

	commando.
		Register("set").
		SetDescription("The set command is used to set a counter's value from the list.").
		SetShortDescription("sets a counter's value").
		AddArgument("name", "Name of the counter", "DefaultCounter").
		AddArgument("value", "Sets the entered value to the counter", "0").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			val, err := strconv.Atoi(args["value"].Value)
			if err != nil {
				log.Fatal("[ERR] Error in setting the counter value.")
			}
			rdb.SetCounter(args["name"].Value, val)
		})

	commando.Parse(nil)
}
