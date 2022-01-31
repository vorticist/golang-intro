package service

import (
	"fmt"
	"reminders/internal/models"
	"reminders/internal/repository"
	"time"
)

func proccess(output models.Output) {
	//Send to write in output database after 5 seconds
	fmt.Println("The reminder will start for test on 5 seconds")
	time.Sleep(5 * time.Second)
	or := repository.NewOutputRepository()
	outputId := or.NewOutput(output)
	if len(outputId) > 0 {
		fmt.Println("Safe data in output data base: ", outputId)
	} else {
		fmt.Println("Cann't write in output data base")
	}
}

func proccessWithDate(dt time.Time, output models.Output) {
	fmt.Println("The reminder will start running on: ", dt.String())
	currentDate := time.Now()
	seconds := currentDate.Sub(dt).Seconds()
	if seconds > 0 {
		fmt.Printf("The reminder will start running on %f seconds \n", seconds/3600)
		time.Sleep(time.Duration(seconds / 3600))
		or := repository.NewOutputRepository()
		output.Description = output.Description + " by Gorutine"
		outputId := or.NewOutput(output)
		if len(outputId) > 0 {
			fmt.Println("Safe data in output data base: ", outputId)
		} else {
			fmt.Println("Cann't write in output data base")
		}
	} else {
		fmt.Println("Could not start because time has passed ")
	}
}

func writeToOutput(ch chan models.Output) {
	or := repository.NewOutputRepository()
	var out models.Output = <-ch	
	if len(out.Emails) > 0 {
		outputId := or.NewOutput(out)
		if len(outputId) > 0 {
			fmt.Println("Safe data in output data base: ", outputId)
		} else {
			fmt.Println("Cann't write in output data base")
		}
	}
}

func proccessWithChannel(dt time.Time, output models.Output) {
	fmt.Println("The reminder will start when read the channel: ")
	currentDate := time.Now()
	seconds := currentDate.Sub(dt.Add(1)).Seconds()
	if seconds > 0 {
		fmt.Printf("The reminder will start running on %f seconds \n", seconds/3600)
		outputChannel := make(chan models.Output)
		defer close(outputChannel)
		go writeToOutput(outputChannel)
		time.Sleep(time.Duration(seconds / 3600))
		output.Description = output.Description + " by Gorutine with Channel"
		outputChannel <- output
	} else {
		fmt.Println("Could not start because time has passed ")
	}
}

func SendReminder(dt time.Time, output models.Output) {
	go proccess(output)
	go proccessWithDate(dt, output)
	proccessWithChannel(dt, output)
}
