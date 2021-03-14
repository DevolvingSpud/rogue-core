//+build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Runs go mod download and then installs the binary.
func Build() error {
	fmt.Println("Starting build...")
	mg.Deps(Tidy, Security)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "install", "./...")
}

func Tidy() error {
	fmt.Println("Running go mod tidy")
	if err := sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}
	return nil
}

func Security() error {
	fmt.Println("Running gosec -no-fail ./...")
	if err := sh.Run("gosec", "-no-fail", "./..."); err != nil {
		return err
	}
	return nil
}
