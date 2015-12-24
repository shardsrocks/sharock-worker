package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func resolve(queue string, args ...interface{}) error {
	if len(args) != 1 {
		fmt.Printf("Invalid arguments length\n")
		return errors.New("Invalid arguments length")
	}

	packageId, ok := args[0].(json.Number)
	if !ok {
		fmt.Printf("Invalid arguments type\n")
		return errors.New("Invalid arguments type")
	}

	packageIdInt, err := packageId.Int64()
	if err != nil {
		fmt.Printf("Invalid arguments type\n")
		return errors.New("Invalid arguments type")
	}

	packageIdStr := strconv.FormatInt(packageIdInt, 10)
	fmt.Printf("Resolve: begin %s\n", packageIdStr)

	cmd := exec.Command("bin/sharock-resolve-worker", packageIdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Printf("Resolve: end %s\n", packageIdStr)

	return nil
}
