package seeds

import (
	"context"
	"errors"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// Seed type
type Seed struct {
	ctx context.Context
}

// Execute will executes the given seed method
func SeedDB(seedMethodNames ...string) error {
	s := Seed{
		ctx: context.Background(),
	}

	seedType := reflect.TypeOf(s)

	// Execute all seeds if no method name is given
	if len(seedMethodNames) == 0 {
		log.Infoln("Running all seeds...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seed method
			err := seed(s, method.Name)
			if err != nil {
				return err
			}
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		err := seed(s, item)
		if err != nil {
			return err
		}
	}

	return nil
}

func seed(s Seed, seedMethodName string) error {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		return errors.New("Seed method not found: " + seedMethodName)
	}
	// Execute the method
	log.Infof("Seeding %s...", seedMethodName)
	result := m.Call(nil)
	if len(result) > 0 {
		if !result[0].IsNil() {
			return result[0].Interface().(error)
		}
	}
	log.Infof("Seed %s successfully executed", seedMethodName)
	return nil
}
