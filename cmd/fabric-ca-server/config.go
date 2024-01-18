package main

import "fabric-admin/util"

const (
	longName = "Hyperledger Fabric Certificate Authority Server"
	cmdName  = "fabric-ca-server"
)

// configInit 설정값 초기화
func (s *ServerCmd) configInit() (err error) {
	if !s.configRequired() {
		return nil
	}

	s.cfgFileName, s.homeDirectory, err = util.ValidateAndReturnAbsConf(s.cfgFileName, s.homeDirectory, cmdName)
	if err != nil {
		return err
	}

	return nil
}
