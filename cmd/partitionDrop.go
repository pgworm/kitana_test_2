package cmd

import (
	"errors"
	"fmt"
	"github.com/msklnko/kitana/db"
	"github.com/msklnko/kitana/partition"
	"github.com/spf13/cobra"
	"os"
	s "strings"
)

var prtDrop = &cobra.Command{
	Use:     "drop",
	Aliases: []string{"rm"},
	Short:   "Drop partition",
	Args: func(cmd *cobra.Command, args []string) error {
		switch l := len(args); l {
		case 0:
			return errors.New("missing arguments (table, partition name)")
		case 1:
			var tables = s.Split(args[0], ".")
			if len(tables) != 2 {
				return errors.New("invalid property, should be schema+table name")
			}
			present, err := db.CheckTablePresent(tables[0], tables[1])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			if !present {
				return errors.New("table " + args[0] + " does not exist")
			}
			return errors.New("partition name is missing")
		default:
			return nil
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var tables = s.Split(args[0], ".")
		err := db.DropPartition(tables[0], tables[1], []string{args[1]})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		show, err := cmd.Flags().GetBool("show")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if show {
			err := partition.PartitionsInfo(tables[0], tables[1])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	},
}
