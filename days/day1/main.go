package day1

func Exec() error {
	if err := ExecA(); err != nil {
		return err
	}

	if err := ExecB(); err != nil {
		return err
	}

	return nil
}
