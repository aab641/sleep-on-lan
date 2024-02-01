package main

func (conf *Configuration) RegisterDefaultCommand() {
	defaultCommand := CommandConfiguration{Operation: "sleep", CommandType: COMMAND_TYPE_EXTERNAL, IsDefault: true, Command: "systemctl suspend"}
	conf.Commands = []CommandConfiguration{defaultCommand}
}

func RegisterPossibleConfigurationFileNames() []PossibleConfigurationFilename {
	var possibleConfigurationFileNames []PossibleConfigurationFilename
	possibleConfigurationFileNames = append(possibleConfigurationFileNames, PossibleConfigurationFilename{"/etc/sol.json", "default configuration filename under /etc/ (linux)"})
	possibleConfigurationFileNames = append(possibleConfigurationFileNames, PossibleConfigurationFilename{"/etc/sleep-on-lan.json", "default configuration filename under /etc/ (linux)"})
	return possibleConfigurationFileNames
}

func ExecuteCommand(Command CommandConfiguration) {
	if Command.CommandType == COMMAND_TYPE_EXTERNAL {
		logger.Infof("Executing operation [" + Command.Operation + "], type [" + Command.CommandType + "], command [" + Command.Command + "]")
		sleepCommandLineImplementation(Command.Command)
	} else {
		logger.Infof("Unknown command type [" + Command.CommandType + "]")
	}
}

func sleepCommandLineImplementation(cmd string) {
	if cmd == "" {
		cmd = "pm-suspend"
	}
	logger.Infof("Sleep implementation [linux], sleep command is [" + cmd + "]")

 	result := Execute(cmd)
	if result.Success {
		logger.Errorf("Command executed successfully:")
		logger.Errorf(result.Output)
	} else {
		logger.Errorf("Command execution failed:")
		logger.Errorf("Error:", result.Error)
		logger.Errorf("Output:", result.Output)
	}
}
