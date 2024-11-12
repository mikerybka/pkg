package dev

func DeployAll() error {
	servers, err := ListServers()
	if err != nil {
		return err
	}
	for _, s := range servers {
		err = Deploy(s)
		if err != nil {
			return err
		}
	}
	return nil
}
